package prof

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getProfHandler(ctx *gin.Context) {
	var Prof Prof
	pid, err := extractProfID(ctx)
	// pid, err := strconv.ParseUint(ctx.Param("pid"), 10, 32)
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
