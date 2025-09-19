package students

import (
	"net/http"

	"backend/db"
	"backend/models"

	"github.com/gin-gonic/gin"
)

// GET /students/available-courses
func GetAvailableCourses(c *gin.Context) {
	uid := c.GetUint("userID")

	// Cari course yang sudah diambil
	var scs []models.StudentCourse
	if err := db.DB.Where("student_id = ?", uid).Find(&scs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch enrollments"})
		return
	}

	courseIDs := make([]uint, 0, len(scs))
	for _, sc := range scs {
		courseIDs = append(courseIDs, sc.CourseID)
	}

	// Ambil course yang belum diambil
	var available []models.Course
	q := db.DB.Preload("Teacher")
	if len(courseIDs) > 0 {
		q = q.Where("id_course NOT IN ?", courseIDs)
	}
	if err := q.Find(&available).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch courses"})
		return
	}

	c.JSON(http.StatusOK, available)
}
