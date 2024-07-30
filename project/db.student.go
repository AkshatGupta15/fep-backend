package project

import "github.com/gin-gonic/gin"

func addSelectedStudent(ctx *gin.Context, selectedStudent *SelectedStudent) error {
	tx := db.WithContext(ctx).Create(selectedStudent)
	return tx.Error
}

func getSelectedStudentForProject(ctx *gin.Context, selectedStudents *[]SelectedStudent, id uint) error {
	tx := db.WithContext(ctx).Where("project_id = ?", id).Find(selectedStudents)
	return tx.Error
}
func getSelectedStudentForStudent(ctx *gin.Context, selectedStudents *[]SelectedStudent, id uint) error {
	tx := db.WithContext(ctx).Where("student_id = ?", id).Find(selectedStudents)
	return tx.Error
}

func deleteSelectedStudent(ctx *gin.Context, id uint) error {
	tx := db.WithContext(ctx).Where("id = ?", id).Delete(&SelectedStudent{})
	return tx.Error
}
