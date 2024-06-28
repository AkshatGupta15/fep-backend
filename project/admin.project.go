package project

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	m "github.com/pclubiitk/fep-backend/middleware"
	"github.com/sirupsen/logrus"
)

func addNewHandler(ctx *gin.Context) {
	var newProjectRequest Project

	if err := ctx.ShouldBindJSON(&newProjectRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := createProject(ctx, &newProjectRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("A new Project %s is added with id %d by %s", newProjectRequest.ProjectName, newProjectRequest.ID, m.GetUserID(ctx))

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully added"})

}
func addNewBulkHandler(ctx *gin.Context) {
	var request []Project

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := createProjects(ctx, &request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("%d projects is added with by %s", len(request), m.GetUserID(ctx))

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully added"})

}
func updateProjectHandler(ctx *gin.Context) {
	var updateProjectRequest Project

	if err := ctx.ShouldBindJSON(&updateProjectRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updateProjectRequest.ProjectID == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Enter Project ID"})
		return
	}
	updated, err := updateProject(ctx, &updateProjectRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !updated {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	logrus.Infof("A project with id %d is updated by %s", updateProjectRequest.ProjectID, m.GetUserID(ctx))

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully updated"})
}
func deleteProjectHandler(ctx *gin.Context) {

	cid, err := strconv.ParseUint(ctx.Param("cid"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = deleteProject(ctx, uint(cid))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("A project with id %d is deleted by %s", cid, m.GetUserID(ctx))

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully deleted"})

}
