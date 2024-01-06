package api_test

import (
	"errors"
	API "nutrition-tracker/pkg/api"
	"reflect"
	"testing"
)

type mockUserRepo struct{}

// Mock one for testing
func (m mockUserRepo) CreateUser(request API.NewUserRequest) (int, error) {
	if request.Name == "test user already created" {
		return 0, errors.New("repository - user already exists in database")
	}
	return 0, nil
}

func TestCreateNewUser(t *testing.T) {
	mockRepo := mockUserRepo{}
	mockUserSRV := API.NewUserService(&mockRepo)

	tests := []struct {
		name    string
		request API.NewUserRequest
		want    error
	}{
		{
			name: "should create a new user successfully",
			request: API.NewUserRequest{
				Name:          "test user",
				WeightGoal:    "maintain",
				Age:           20,
				Height:        180,
				Sex:           "female",
				ActivityLevel: 5,
				Email:         "test_user@gmail.com",
			},
			want: nil,
		}, {
			name: "should return an error because of missing email",
			request: API.NewUserRequest{
				Name:          "test user",
				Age:           20,
				WeightGoal:    "maintain",
				Height:        180,
				Sex:           "female",
				ActivityLevel: 5,
				Email:         "",
			},
			want: errors.New("user service - email required"),
		}, {
			name: "should return an error because of missing name",
			request: API.NewUserRequest{
				Name:          "",
				Age:           20,
				WeightGoal:    "maintain",
				Height:        180,
				Sex:           "female",
				ActivityLevel: 5,
				Email:         "test_user@gmail.com",
			},
			want: errors.New("user service - name required"),
		}, {
			name: "should return error from database because user already exists",
			request: API.NewUserRequest{
				Name:          "test user already created",
				Age:           20,
				Height:        180,
				WeightGoal:    "maintain",
				Sex:           "female",
				ActivityLevel: 5,
				Email:         "test_user@gmail.com",
			},
			want: errors.New("repository - user already exists in database"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := mockUserSRV.New(test.request)

			if !reflect.DeepEqual(err, test.want) {
				t.Errorf("test: %v failed. Got: %v, wanted: %v", test.name, err, test.want)
			}
		})
	}
}
