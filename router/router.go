package router

import (
	"GinGormCRUD/controller"
	"GinGormCRUD/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// Simple group: v1
	v1 := r.Group("/v1")
	{
		login := v1.Group("/login")
		{
			login.POST("/login", controller.Login)
			login.POST("/register", controller.Register)
		}

		seat := v1.Group("/seat")
		seat.GET("/getSeatList", controller.SeatsList)
		seat.Use(middleware.AuthMiddleware())
		{
			seat.POST("/postSeat", controller.SetSeat)
		}

	}

	return r
}
