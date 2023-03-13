package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pablouser1/AreYaCoding/helpers/tokens"
	"github.com/pablouser1/AreYaCoding/models"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := tokens.Verify(c.GetHeader("Authorization"))
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.ErrorRes{
				Code: http.StatusUnauthorized,
				Msg:  "Unauthorized",
			})
			c.Abort()
			return
		}

		c.Set("username", username)

		c.Next()
	}
}
