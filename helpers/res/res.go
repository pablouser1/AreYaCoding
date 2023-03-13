package res

import (
	"github.com/gin-gonic/gin"
	"github.com/pablouser1/AreYaCoding/models"
)

func Send(c *gin.Context, status int, data interface{}, args ...string) {
	msg := "OK"
	success := status >= 200 && status < 400

	if len(args) > 0 {
		msg = args[0]
	}

	c.JSON(status, models.ApiRes{
		Status:  status,
		Success: success,
		Msg:     msg,
		Data:    data,
	})
}
