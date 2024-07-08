package project

import (
	"github.com/gin-gonic/gin"
)

func AdminRouter(r *gin.Engine) {
	admin := r.Group("/api/admin")
	{
		admin.GET("/project", getAllProjectsHandler)
		admin.GET("/project/:cid", getProjectHandler)
		admin.GET("/project/limited", getLimitedProjectsHandler)

		admin.PUT("/project", updateProjectHandler)
		admin.POST("/project", addNewHandler)
		admin.POST("/project/bulk", addNewBulkHandler)

		admin.DELETE("/project/:cid", deleteProjectHandler)

		admin.GET("/application", getAllApplicationsHandler)
		admin.GET("/application/project/:cid", getApplicationsForProjectHandler)
		admin.GET("/application/student/:cid", getApplicationsofStudentHandler)

		admin.PUT("/application", updateApplicationHandler)

		admin.DELETE("/application/:cid", deleteApplicationHandler)

	}

}
func StudentRouter(r *gin.Engine) {
	student := r.Group("/api/student/application")
	{
		student.POST("", addApplicationHandler)
		student.GET("/:cid", getApplicationsofStudentHandler)
		student.PUT("", updateApplicationHandler)
		student.DELETE("/:cid", deleteApplicationHandler)
	}
}
