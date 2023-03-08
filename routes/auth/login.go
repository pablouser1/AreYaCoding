package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pablouser1/AreYaCoding/helpers/tokens"
	"github.com/pablouser1/AreYaCoding/models"
)

func Login(c *gin.Context) {
	username := c.Query("username")
	if username != "" {
		token, err := tokens.Create(username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorRes{
				Code: http.StatusInternalServerError,
				Msg:  err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, models.Auth{
			Token: token,
		})
	}
}
