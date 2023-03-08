package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pablouser1/AreYaCoding/helpers/tokens"
)

func Test(c *gin.Context) {
	username, _ := tokens.Verify(c.GetHeader("Authorization"))
	c.String(http.StatusOK, username)
}
