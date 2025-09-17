package courses

import (
	"backend/db"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateCourse(c *gin.Context) {
	var course models.Course
	if err := db.DB.First(&course, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}
	db.DB.Save(&course)
	c.JSON(http.StatusOK, course)
}
