package rooms

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/pablouser1/AreYaCoding/models"
)

func Create(c *gin.Context) {
	roomName := c.Query("roomName")
	if roomName != "" {
		room := models.Room{
			Name: roomName,
			Slug: slug.Make(roomName),
		}

		models.DB.Create(&room)
		c.JSON(http.StatusOK, room)
	} else {
		c.JSON(http.StatusBadRequest, models.ErrorRes{
			Code: http.StatusBadRequest,
			Msg:  "Invalid request",
		})
	}
}
