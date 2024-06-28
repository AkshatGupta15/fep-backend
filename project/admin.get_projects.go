package project

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pclubiitk/fep-backend/util"
)

func getAllProjectsHandler(ctx *gin.Context) {
	var projects []Project

	err := getAllProjects(ctx, &projects)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, projects)
}
func getProjectHandler(ctx *gin.Context) {
	var project Project

	cid, err := strconv.ParseUint(ctx.Param("cid"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = getProject(ctx, &project, uint(cid))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, project)
}
func getLimitedProjectsHandler(ctx *gin.Context) {
	var projects []Project

	pageSize := ctx.DefaultQuery("pageSize", "100")
	lastFetchedId := ctx.Query("lastFetchedId")
	pageSizeInt, err := util.ParseUint(pageSize)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lastFetchedIdInt, err := util.ParseUint(lastFetchedId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = getLimitedProjects(ctx, &projects, uint(lastFetchedIdInt), int(pageSizeInt))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, projects)
}
