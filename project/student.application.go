package project

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	m "github.com/pclubiitk/fep-backend/middleware"
	"github.com/sirupsen/logrus"
)

func addApplicationHandler(ctx *gin.Context) {
	var application Application

	if err := ctx.ShouldBindJSON(&application); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := createApplication(ctx, &application)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("A new Application %s is added with id %d by %s", application.ProjectID, application.ID, m.GetUserID(ctx))

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully added"})

}
func updateApplicationHandler(ctx *gin.Context) {
	var updateApplicationRequest Application

	if err := ctx.ShouldBindJSON(&updateApplicationRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updateApplicationRequest.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Enter Application ID"})
		return
	}
	updated, err := updateApplicationByStudent(ctx, &updateApplicationRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !updated {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Application not found"})
		return
	}

	logrus.Infof("A Application with id %d is updated by %s", updateApplicationRequest.ID, m.GetUserID(ctx))

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully updated"})
}
func deleteApplicationHandler(ctx *gin.Context) {

	cid, err := strconv.ParseUint(ctx.Param("cid"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = deleteApplication(ctx, uint(cid))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("A Application with id %d is deleted by %s", cid, m.GetUserID(ctx))

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully deleted"})

}
