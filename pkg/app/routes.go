package app

import "github.com/gin-gonic/gin"

func (s *Server) Routes() *gin.Engine {
	router := s.router

	// Group all routes under /v1/api
	v1 := router.Group("/v1/api")
	{
		v1.GET("/status", s.APIStatus())

		user := v1.Group("/user")
		{
			user.POST("", s.CreateUser())
		}

		weight := v1.Group("/weight")
		{
			weight.POST("", s.CreateWeightEntry())
		}
	}

	return router
}
