package students

import (
	"net/http"

	"backend/db"
	"backend/models"

	"github.com/gin-gonic/gin"
)

type enrollReq struct {
	CourseID uint `json:"course_id" binding:"required"`
}

func PostEnroll(c *gin.Context) {
	uid := c.GetUint("userID")

	var req enrollReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// cek apakah course ada
	var course models.Course
	if err := db.DB.First(&course, req.CourseID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "course not found"})
		return
	}

	sc := models.StudentCourse{
		StudentID: uid,
		CourseID:  req.CourseID,
	}
	if err := db.DB.Create(&sc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "already enrolled or invalid"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "enrolled"})
}
