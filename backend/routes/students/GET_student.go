package students

import (
	"backend/db"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	var students []models.Student
	db.DB.Preload("Courses").Find(&students)
	c.JSON(http.StatusOK, students)
}
