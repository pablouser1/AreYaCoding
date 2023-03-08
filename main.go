package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pablouser1/AreYaCoding/helpers"
	"github.com/pablouser1/AreYaCoding/middlewares"
	"github.com/pablouser1/AreYaCoding/models"
	"github.com/pablouser1/AreYaCoding/routes/auth"
	"github.com/pablouser1/AreYaCoding/routes/rooms"
	"github.com/pablouser1/AreYaCoding/routes/ws"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not load .env ->", err.Error())
		os.Exit(1)
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

	port := fmt.Sprintf(":%s", helpers.GetEnv("PORT", "8080"))

	err = r.Run(port)

	if err != nil {
		fmt.Println(err)
	}
}
