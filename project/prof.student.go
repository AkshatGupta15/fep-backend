package project

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	m "github.com/pclubiitk/fep-backend/middleware"
	"github.com/sirupsen/logrus"
)

func addSelectedStudentHandler(ctx *gin.Context) {
	var newStudentRequest SelectedStudent

	if err := ctx.ShouldBindJSON(&newStudentRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := addSelectedStudent(ctx, &newStudentRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("A new student is selected with id %d for project %d by %s", newStudentRequest.StudentID, newStudentRequest.ProjectID, m.GetUserID(ctx))

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully added"})

}
func deleteSelectedStudentHandler(ctx *gin.Context) {

	cid, err := strconv.ParseUint(ctx.Param("cid"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = deleteSelectedStudent(ctx, uint(cid))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("A selected student with id %d is deleted by %s", cid, m.GetUserID(ctx))

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully deleted"})

}
func getSelectedStudentForProjectHandler(ctx *gin.Context) {
	var selectedStudents []SelectedStudent

	cid, err := strconv.ParseUint(ctx.Param("cid"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = getSelectedStudentForProject(ctx, &selectedStudents, uint(cid))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, selectedStudents)
}
func getSelectedStudentForStudentHandler(ctx *gin.Context) {
	var selectedStudents []SelectedStudent

	cid, err := strconv.ParseUint(ctx.Param("cid"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = getSelectedStudentForStudent(ctx, &selectedStudents, uint(cid))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, selectedStudents)
}
