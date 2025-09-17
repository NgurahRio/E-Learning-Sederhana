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

type enrollResp struct {
	Message     string `json:"message"`
	IDCourse    uint   `json:"id_course"`
	TeacherName string `json:"teacher_name"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func PostEnroll(c *gin.Context) {
	uid := c.GetUint("userID")

	var req enrollReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// cek apakah course ada + preload teacher
	var course models.Course
	if err := db.DB.Preload("Teacher").First(&course, req.CourseID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "pelajaran tidak ditemukan"})
		return
	}

	// cek apakah student sudah terdaftar di course ini
	var existing models.StudentCourse
	if err := db.DB.Where("student_id = ? AND course_id = ?", uid, req.CourseID).First(&existing).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Anda sudah terdaftar di pelajaran ini"})
		return
	}

	// buat relasi student-course baru
	sc := models.StudentCourse{
		StudentID: uid,
		CourseID:  req.CourseID,
	}
	if err := db.DB.Create(&sc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to enroll"})
		return
	}

	// response sukses
	resp := enrollResp{
		Message:     "Berhasil mendaftar di pelajaran",
		IDCourse:    course.IDCourse,
		TeacherName: course.Teacher.Name,
		Title:       course.Title,
		Description: course.Description,
	}

	c.JSON(http.StatusCreated, resp)
}
