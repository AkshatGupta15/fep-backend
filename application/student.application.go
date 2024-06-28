package application

import (
	"net/http"
	

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

	logrus.Infof("A new Application %s is added with id %d by %s", application.ProjectName, application.ID, m.GetUserID(ctx))

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully added"})

}