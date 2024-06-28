package application

import "github.com/gin-gonic/gin"

func createApplication(ctx *gin.Context, application*Application) error {
	tx := db.WithContext(ctx).Create(application)
	return tx.Error
}