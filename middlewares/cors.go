package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/pablouser1/AreYaCoding/helpers"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Header("Access-Control-Max-Age", "86400")
		c.Header("Access-Control-Allow-Origin", helpers.GetEnv("APP_URL", "http://localhost:3002"))
		if c.Request.Method == "OPTIONS" {
			c.Abort()
		}
		c.Next()
	}
}
