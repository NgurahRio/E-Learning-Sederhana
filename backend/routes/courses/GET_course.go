package courses

import (
	"backend/db"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCourse(c *gin.Context) {
	var courses []models.Course
	db.DB.Preload("Students").Find(&courses)
	c.JSON(http.StatusOK, courses)
}
