package project

import (
	"github.com/gin-gonic/gin"
)

func AdminRouter(r *gin.Engine) {
	admin := r.Group("/api/admin/project")
	{
		admin.GET("", getAllProjectsHandler)
		admin.GET("/:cid", getProjectHandler)
		admin.GET("/limited", getLimitedProjectsHandler)

		admin.PUT("", updateProjectHandler)
		admin.POST("", addNewHandler)
		admin.POST("/bulk", addNewBulkHandler)

		admin.DELETE("/:cid", deleteProjectHandler)

	}
}
func StudentRouter(r *gin.Engine) {
    student := r.Group("/api/student/application")
    {
        student.POST("", addApplicationHandler)  // Route for students to add an application
    }
}