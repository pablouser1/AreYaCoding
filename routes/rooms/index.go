package rooms

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pablouser1/AreYaCoding/models"
)

func GetAll(c *gin.Context) {
	var rooms []models.Room
	models.DB.Find(&rooms)

	c.JSON(http.StatusOK, rooms)
}
