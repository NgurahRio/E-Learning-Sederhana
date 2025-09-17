package main

import (
	"backend/db"
	"backend/enroll"
	"backend/routes/courses"
	"backend/routes/students"
	"backend/routes/users"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Users
	r.GET("/api/users", users.GetUsers)
	r.POST("/api/users", users.CreateUsers)
	r.PUT("/api/users/:id", users.UpdateUsers)
	r.DELETE("/api/users/:id", users.DeleteUsers)

	// Students
	r.GET("/api/students", students.GetStudents)
	r.POST("/api/students", students.CreateStudents)
	r.PUT("/api/students/:id", students.UpdateStudents)
	r.DELETE("/api/students/:id", students.DeleteStudents)

	// Courses
	r.GET("/api/courses", courses.GetCourse)
	r.POST("/api/courses", courses.CreateCourse)
	r.PUT("/api/courses/:id", courses.UpdateCourse)
	r.DELETE("/api/courses/:id", courses.DeleteCourse)

	// Enroll
	r.POST("/api/enroll", enroll.EnrollStudent)

	r.Run(":8080")
}
