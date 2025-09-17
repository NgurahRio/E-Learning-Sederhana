package users

import (
	"net/http"

	"backend/db"
	"backend/models"

	"github.com/gin-gonic/gin"
)

type updateUserReq struct {
	Name   *string `json:"name"`
	RoleID *uint   `json:"role_id"`
}

func PutUser(c *gin.Context) {
	id := c.Param("id")

	// cari user
	var user models.User
	if err := db.DB.First(&user, "id_user = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	// bind request
	var req updateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update field jika ada
	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.RoleID != nil {
		user.RoleID = *req.RoleID
	}

	// simpan ke DB
	if err := db.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		return
	}

	// response tanpa password
	c.JSON(http.StatusOK, gin.H{
		"id":      user.IDUser,
		"name":    user.Name,
		"email":   user.Email,
		"role_id": user.RoleID,
	})
}
