package courses

import (
	"backend/db"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteCourse(c *gin.Context) {
	if err := db.DB.Delete(&models.Course{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete"})
		return
	}
	c.Status(http.StatusNoContent)
}
