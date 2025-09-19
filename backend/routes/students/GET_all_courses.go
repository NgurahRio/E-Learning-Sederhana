package students

import (
	"net/http"

	"backend/db"
	"backend/models"

	"github.com/gin-gonic/gin"
)

// GET /students/all-courses
func GetAllCourses(c *gin.Context) {
	var courses []models.Course

	// preload teacher kalau ada relasi ke users
	if err := db.DB.Preload("Teacher").Find(&courses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch courses"})
		return
	}

	c.JSON(http.StatusOK, courses)
}
