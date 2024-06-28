package application

import (
    "github.com/gin-gonic/gin"
)

func StudentRouter(r *gin.Engine) {
    student := r.Group("/api/student/application")
    {
        student.POST("", addApplicationHandler)  // Route for students to add an application
    }
}
