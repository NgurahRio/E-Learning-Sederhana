package teachers

import (
	"net/http"

	"backend/db"
	"backend/models"

	"github.com/gin-gonic/gin"
)

type courseReq struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

func PostCourse(c *gin.Context) {
	uid := c.GetUint("userID")

	var req courseReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course := models.Course{
		Title:       req.Title,
		Description: req.Description,
		TeacherID:   uid,
	}

	if err := db.DB.Create(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create"})
		return
	}
	c.JSON(http.StatusCreated, course)
}
