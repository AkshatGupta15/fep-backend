package project

import "github.com/gin-gonic/gin"

func getProject(ctx *gin.Context, project *Project, id uint) error {
	tx := db.WithContext(ctx).Where("id = ?", id).First(project)
	return tx.Error
}
func getAllProjects(ctx *gin.Context, projects *[]Project) error {
	tx := db.WithContext(ctx).Find(projects)
	return tx.Error
}
func getLimitedProjects(ctx *gin.Context, projects *[]Project, lastFetchedId uint, pageSize int) error {
	tx := db.WithContext(ctx).Order("id asc").Where("id >= ?", lastFetchedId).Limit(pageSize).Find(projects)
	return tx.Error
}
func updateProject(ctx *gin.Context, project *Project) (bool, error) {
	tx := db.WithContext(ctx).Where("id = ?", project.ID).Updates(project)
	return tx.RowsAffected > 0, tx.Error
}
func createProject(ctx *gin.Context, project *Project) error {
	tx := db.WithContext(ctx).Create(project)
	return tx.Error
}
func createProjects(ctx *gin.Context, projects *[]Project) error {
	tx := db.WithContext(ctx).Create(projects)
	return tx.Error
}
func deleteProject(ctx *gin.Context, id uint) error {
	tx := db.WithContext(ctx).Where("id = ?", id).Delete(&Project{})
	return tx.Error
}
func getProjectsFromProf(ctx *gin.Context, projects *[]Project, id uint) error {
	tx := db.WithContext(ctx).Where("prof_id = ?", id).Find(projects)
	return tx.Error
}
func GetProjectName(ctx *gin.Context, id uint) (string, error) {
	var p Project
	err := getProject(ctx, &p, id)
	if err != nil {
		return "", err
	}
	return p.ProjectName, nil
}

func GetUniversityName(ctx *gin.Context, id uint) (string, error) {
	var p Project
	err := getProject(ctx, &p, id)
	if err != nil {
		return "", err
	}
	return p.UniversityName, nil
}
