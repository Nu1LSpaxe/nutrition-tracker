package repository

import (
	"database/sql"
	"errors"
	"log"
	"path/filepath"
	"runtime"

	API "nutrition-tracker/pkg/api"

	"github.com/golang-migrate/migrate/v4" // Database migrations written in Go. Use as CLI.
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq" // Go postgres driver
)

type Storage interface {
	RunMigrations(conn string) error
	CreateUser(request API.NewUserRequest) (int, error)
	CreateWeightEntry(request API.Weight) error
	GetUser(userID int) (API.User, error)
}

type storage struct {
	DB *sql.DB
}

func NewStorage(DB *sql.DB) Storage {
	return &storage{
		DB: DB,
	}
}

func (s *storage) RunMigrations(DBURI string) error {
	if DBURI == "" {
		return errors.New("repository: the DBURI was empty")
	}

	// Get base path
	_, file, _, _ := runtime.Caller(0)
	basePath := filepath.Join(filepath.Dir(file), "../..")

	migrationsPath := "file://" + filepath.ToSlash(basePath) + "/pkg/repository/migrations/"

	m, err := migrate.New(migrationsPath, DBURI)
	if err != nil {
		log.Println("wrong with migrations")
		return err
	}

	err = m.Up()

	switch err {
	case errors.New("no cahnge"):
		return nil
	}

	return nil
}

func (s *storage) CreateUser(request API.NewUserRequest) (int, error) {
	newUserStatement := `
		INSERT INTO "user" (name, age, height, sex, activity_level, email, weight_goal)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id;
		`

	userID := 0
	err := s.DB.QueryRow(newUserStatement, request.Name, request.Age, request.Height, request.Sex, request.ActivityLevel, request.Email, request.WeightGoal).Scan(&userID)

	if err != nil {
		log.Printf("Error occurs when creating user: %v", err.Error())
		return 0, err
	}

	return userID, nil
}

func (s *storage) CreateWeightEntry(request API.Weight) error {
	newWeight := `
		INSERT INTO weight (weight, user_id, bmr, daily_caloric_intake)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
		`

	var ID int
	err := s.DB.QueryRow(newWeight, request.Weight, request.UserID, request.BMR, request.DailyCaloricIntake).Scan(&ID)

	if err != nil {
		log.Printf("Error occurs when creating weight entry: %v", err.Error())
		return err
	}

	return nil
}

func (s *storage) GetUser(userID int) (API.User, error) {
	getUser := `
		SELECT id, name, age, height, sex, activity_level, email, weight_goal FROM "user"
		where id=$1;
		`

	var user API.User
	err := s.DB.QueryRow(getUser, userID).Scan(&user)

	if err != nil {
		log.Printf("Error occurs when getting user: %v", err.Error())
		return API.User{}, err
	}

	return user, nil
}
