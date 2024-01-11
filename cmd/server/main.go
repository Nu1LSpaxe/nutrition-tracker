package main

import (
	"database/sql"
	"fmt"
	"log"
	API "nutrition-tracker/pkg/api"
	APP "nutrition-tracker/pkg/app"
	"nutrition-tracker/pkg/repository"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "start error with %v\n", err)
		os.Exit(1)
	}
}

// Setting up database connections, routers, etc.
func run() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error occurs whan loading environment file.")
	}

	DBURI := os.Getenv("PostgreSQL_URI")
	if DBURI == "" {
		log.Fatal("You haven't set up Database URI.")
	}

	// Set up database connection
	DB, err := connectDB(DBURI)
	if err != nil {
		return err
	}

	// Create storage dependency
	storage := repository.NewStorage(DB)
	// Run migrations
	err = storage.RunMigrations(DBURI)
	if err != nil {
		return err
	}

	// Create router dependency
	router := gin.Default()
	router.Use(cors.Default())

	// Create user service
	userSRV := API.NewUserService(storage)

	// Create weight service
	weightSRV := API.NewWeightService(storage)

	server := APP.NewServer(router, userSRV, weightSRV)

	// Run server
	err = server.Run()
	if err != nil {
		return err
	}

	return nil
}

func connectDB(DBURI string) (*sql.DB, error) {
	DB, err := sql.Open("postgres", DBURI)
	if err != nil {
		return nil, err
	}

	// Ping DB to ensure connection
	err = DB.Ping()
	if err != nil {
		return nil, err
	}

	return DB, err
}
