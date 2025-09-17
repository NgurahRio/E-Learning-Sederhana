package students

import (
	"backend/db"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateStudents(c *gin.Context) {
	var student models.Student
	if err := db.DB.First(&student, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}
	db.DB.Save(&student)
	c.JSON(http.StatusOK, student)
}
