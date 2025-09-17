package main

import (
	"log"
	"os"

	"backend/db"
	"backend/middleware"
	"backend/routes/students"
	"backend/routes/teachers"
	"backend/routes/users"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	r := gin.Default()

	// Public routes
	pub := r.Group("/api")
	{
		pub.POST("/users/register", users.PostUser)
		pub.POST("/users/login", users.PostLogin)
	}

	// Protected routes
	auth := r.Group("/api", middleware.JWTAuth())
	{
		// Users (admin only)
		auth.GET("/users", middleware.RequireRoles("admin"), users.GetUserList)
		auth.PUT("/users/:id", middleware.RequireRoles("admin"), users.PutUser)
		auth.DELETE("/users/:id", middleware.RequireRoles("admin"), users.DeleteUser)

		// Students
		stu := auth.Group("/students", middleware.RequireRoles("student"))
		{
			stu.GET("/courses", students.GetStudentCourses)
			stu.POST("/enroll", students.PostEnroll)
			// stu.PUT("/profile", students.PutStudentProfile)
		}

		// Teachers
		tch := auth.Group("/teachers", middleware.RequireRoles("teacher"))
		{
			tch.GET("/my-courses", teachers.GetMyCourses)
			tch.POST("/course", teachers.PostCourse)
			tch.PUT("/course/:id", teachers.PutCourse)
			tch.DELETE("/course/:id", teachers.DeleteCourse)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("✅ Server running on :" + port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
