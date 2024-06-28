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
