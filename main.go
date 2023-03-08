package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pablouser1/AreYaCoding/middlewares"
	"github.com/pablouser1/AreYaCoding/models"
	"github.com/pablouser1/AreYaCoding/routes/auth"
	"github.com/pablouser1/AreYaCoding/routes/rooms"
	"github.com/pablouser1/AreYaCoding/routes/ws"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		return
	}

	models.ConnectDatabase()
	hub := ws.NewHub()
	go hub.Run()

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	v1 := r.Group("/api/v1")
	{
		authG := v1.Group("/auth")
		{
			authG.POST("/login", auth.Login)
			authG.GET("/test", auth.Test)
		}

		roomsG := v1.Group("/rooms")
		{
			roomsG.GET("/", rooms.GetAll)
			roomsG.POST("/", rooms.Create)
			roomsG.GET("/:id", rooms.One)
			roomsG.GET("/:id/ws", middlewares.JwtAuthMiddleware(), func(c *gin.Context) {
				roomId := c.Param("id")
				ws.ServeWS(c, roomId, hub)
			})
		}
	}

	r.Run("localhost:3001")
}
