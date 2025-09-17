package enroll

import (
	"backend/db"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EnrollStudent(c *gin.Context) {
	var body struct {
		StudentID uint `json:"student_id"`
		CourseID  uint `json:"course_id"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}

	link := models.StudentCourse{
		StudentID: body.StudentID,
		CourseID:  body.CourseID,
	}
	if err := db.DB.Create(&link).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Student enrolled to course"})
}
