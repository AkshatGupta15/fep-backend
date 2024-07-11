package prof

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	m "github.com/pclubiitk/fep-backend/middleware"
	"github.com/sirupsen/logrus"
)

func greetingHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Hey!!"})
}
func addNewHandler(ctx *gin.Context) {
	var newProfRequest Prof

	if err := ctx.ShouldBindJSON(&newProfRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := createProf(ctx, &newProfRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("A new Prof %s is added with id %d by %s", newProfRequest.ProfessorName, newProfRequest.ID, m.GetUserID(ctx))

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully added"})

}

// func addNewBulkHandler(ctx *gin.Context) {
// 	var request []Company

// 	if err := ctx.ShouldBindJSON(&request); err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	err := createCompanies(ctx, &request)
// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	logrus.Infof("%d companies is added with by %s", len(request), m.GetUserID(ctx))

// 	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully added"})

// }

func updateProfHandler(ctx *gin.Context) {
	var updateCompanyRequest Prof

	if err := ctx.ShouldBindJSON(&updateCompanyRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updateCompanyRequest.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Enter Prof ID"})
		return
	}
	updated, err := updateProf(ctx, &updateCompanyRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !updated {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Prof not found"})
		return
	}

	logrus.Infof("A prof with id %d is updated by %s", updateCompanyRequest.ID, m.GetUserID(ctx))

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully updated"})
}

func deleteProfHandler(ctx *gin.Context) {

	cid, err := strconv.ParseUint(ctx.Param("cid"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = deleteProf(ctx, uint(cid))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("A Prof with id %d is deleted by %s", cid, m.GetUserID(ctx))

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully deleted"})

}
