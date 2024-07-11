package prof

import "gorm.io/gorm"

type Prof struct {
	gorm.Model
	ProfessorName    string `json:"professor_name"`
	ProfessorEmailId string `json:"professor_email_id"`
	UniversityName   string `json:"university_name"`
}
