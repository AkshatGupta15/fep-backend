package prof

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pclubiitk/fep-backend/util"
)

func getAllProfHandler(ctx *gin.Context) {
	var companies []Prof

	err := getAllProf(ctx, &companies)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, companies)
}

func getProfHandlerAdmin(ctx *gin.Context) {
	var Prof Prof
	// cid, err := extractProfID(ctx)
	pid, err := strconv.ParseUint(ctx.Param("pid"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = getProf(ctx, &Prof, uint(pid))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, Prof)
}

func getLimitedProfHandler(ctx *gin.Context) {
	var companies []Prof

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
	err = getLimitedProf(ctx, &companies, uint(lastFetchedIdInt), int(pageSizeInt))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, companies)
}
