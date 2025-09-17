package users

import (
    "net/http"
    "backend/db"
    "backend/models"
    "github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
    var users []models.User
    db.DB.Find(&users)
    c.JSON(http.StatusOK, users)
}
