package students

import (
	"net/http"

	"backend/db"
	"backend/models"

	"github.com/gin-gonic/gin"
)

// response khusus untuk student courses
type studentCourseResp struct {
	IDCourse    uint   `json:"id_course"`
	TeacherName string `json:"teacher_name"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func GetStudentCourses(c *gin.Context) {
	uid := c.GetUint("userID")

	var scs []models.StudentCourse
	if err := db.DB.Where("student_id = ?", uid).Find(&scs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch"})
		return
	}

	// ambil semua course_id
	courseIDs := make([]uint, 0, len(scs))
	for _, sc := range scs {
		courseIDs = append(courseIDs, sc.CourseID)
	}

	var courses []models.Course
	if len(courseIDs) > 0 {
		if err := db.DB.Preload("Teacher").Find(&courses, courseIDs).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch courses"})
			return
		}
	}

	// mapping ke response
	resps := make([]studentCourseResp, 0, len(courses))
	for _, ccourse := range courses {
		resps = append(resps, studentCourseResp{
			IDCourse:    ccourse.IDCourse,
			TeacherName: ccourse.Teacher.Name,
			Title:       ccourse.Title,
			Description: ccourse.Description,
		})
	}

	c.JSON(http.StatusOK, resps)
}
