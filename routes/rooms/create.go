package rooms

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/pablouser1/AreYaCoding/helpers/res"
	"github.com/pablouser1/AreYaCoding/models"
)

func Create(c *gin.Context) {
	roomName := c.Query("roomName")
	if roomName == "" {
		res.Send(c, http.StatusBadRequest, gin.H{}, "You need to send a room name")
	}

	room := models.Room{
		Name: roomName,
		Slug: slug.Make(roomName),
	}

	models.DB.Create(&room)
	res.Send(c, http.StatusOK, room)
}
