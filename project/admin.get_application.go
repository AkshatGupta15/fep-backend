package project

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getAllApplicationsHandler(ctx *gin.Context) {
	var applications []Application

	err := getAllApplications(ctx, &applications)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, applications)
}
func getApplicationsForProjectHandler(ctx *gin.Context) {
	var applications []Application

	cid, err := strconv.ParseUint(ctx.Param("cid"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = getApplicationsForProject(ctx, &applications, uint(cid))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, applications)
}
func getApplicationsofStudentHandler(ctx *gin.Context) {
	var applications []Application

	cid, err := strconv.ParseUint(ctx.Param("cid"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = getApplicationsForStudent(ctx, &applications, uint(cid))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, applications)
}
