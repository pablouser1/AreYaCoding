package rooms

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pablouser1/AreYaCoding/models"
)

func One(c *gin.Context) {
	var room models.Room

	res := models.DB.Where("slug = ?", c.Param("id"))

	if res.Error == nil {
		res.Find(&room)
		if res.RowsAffected > 0 {
			c.JSON(http.StatusOK, room)
		} else {
			c.JSON(http.StatusNotFound, models.ErrorRes{
				Code: http.StatusNotFound,
				Msg:  "Room not found",
			})
		}
	} else {
		c.JSON(http.StatusInternalServerError, models.ErrorRes{
			Code: http.StatusInternalServerError,
			Msg:  res.Error.Error(),
		})
	}
}
