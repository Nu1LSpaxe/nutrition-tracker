package app

import (
	"log"
	"net/http"
	API "nutrition-tracker/pkg/api"

	"github.com/gin-gonic/gin"
)

func (s *Server) APIStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		res := map[string]string{
			"status": "success",
			"data":   "weight tracker: running smoothly",
		}

		c.JSON(http.StatusOK, res)
	}
}

func (s *Server) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var err error
		var userID int
		var newUser API.NewUserRequest

		err = c.ShouldBindJSON(&newUser)

		if err != nil {
			log.Printf("handler error: %v\n", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		userID, err = s.userSRV.New(newUser)

		if err != nil {
			log.Printf("service error: %v\n", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		log.Printf("new user(id: %v) created\n", userID)

		res := map[string]string{
			"status": "success",
			"data":   "new user created",
		}

		c.JSON(http.StatusOK, res)
	}
}

func (s *Server) CreateWeightEntry() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var newWeight API.NewWeightRequest
		err := c.ShouldBindJSON(&newWeight)

		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		err = s.weightSRV.New(newWeight)

		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		response := map[string]string{
			"status": "success",
			"data":   "new weight created",
		}

		c.JSON(http.StatusOK, response)
	}
}
