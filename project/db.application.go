package project

import "github.com/gin-gonic/gin"

func createApplication(ctx *gin.Context, application *Application) error {
	tx := db.WithContext(ctx).Create(application)
	return tx.Error
}
func getApplicationsForProject(ctx *gin.Context, applications *[]Application, id uint) error {
	tx := db.WithContext(ctx).Where("project_id = ?", id).Find(applications)
	return tx.Error
}
func getAllApplications(ctx *gin.Context, applications *[]Application) error {
	tx := db.WithContext(ctx).Find(applications)
	return tx.Error
}
func updateApplicationByStudent(ctx *gin.Context, application *Application) (bool, error) {
	tx := db.WithContext(ctx).Where("id = ?", application.ID).Updates(application)
	return tx.RowsAffected > 0, tx.Error
}
func getApplicationsForStudent(ctx *gin.Context, applications *[]Application, id uint) error {
	tx := db.WithContext(ctx).Where("student_id = ?", id).Find(applications)
	return tx.Error
}
func deleteApplication(ctx *gin.Context, id uint) error {
	tx := db.WithContext(ctx).Where("id = ?", id).Delete(&Application{})
	return tx.Error
}
