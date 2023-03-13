package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pablouser1/AreYaCoding/helpers/res"
	"github.com/pablouser1/AreYaCoding/helpers/tokens"
	"github.com/pablouser1/AreYaCoding/models"
)

func Login(c *gin.Context) {
	username := c.Query("username")
	if username != "" {
		token, err := tokens.Create(username)
		if err != nil {
			res.Send(c, http.StatusInternalServerError, gin.H{}, err.Error())
			return
		}

		res.Send(c, http.StatusOK, models.Auth{
			Token: token,
		})
	}
}
