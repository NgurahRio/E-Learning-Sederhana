package teachers

import (
	"net/http"

	"backend/db"
	"backend/models"

	"github.com/gin-gonic/gin"
)

func GetMyCourses(c *gin.Context) {
	uid := c.GetUint("userID")

	var courses []models.Course
	if err := db.DB.Where("teacher_id = ?", uid).Find(&courses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch"})
		return
	}
	c.JSON(http.StatusOK, courses)
}
