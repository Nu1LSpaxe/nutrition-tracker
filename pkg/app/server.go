package app

import (
	"log"
	API "nutrition-tracker/pkg/api"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router    *gin.Engine
	userSRV   API.UserService
	weightSRV API.WeightService
}

func NewServer(router *gin.Engine, userService API.UserService, weightService API.WeightService) *Server {
	return &Server{
		router:    router,
		userSRV:   userService,
		weightSRV: weightService,
	}
}

func (s *Server) Run() error {
	// Initializes the routes
	r := s.Routes()

	// Run server through the router
	err := r.Run()

	if err != nil {
		log.Printf("Error occurs when calling Server.Run() on router: %v", err)
		return err
	}

	return nil
}
