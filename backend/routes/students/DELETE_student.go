package students

import (
	"backend/db"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteStudents(c *gin.Context) {
	if err := db.DB.Delete(&models.Student{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete"})
		return
	}
	c.Status(http.StatusNoContent)
}
