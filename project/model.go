package project

import (
	"time"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	ProjectID           uint      `json:"project_id"`
	ProjectName         string    `json:"project_name"`
	ProfessorName       string    `json:"professor_name"`
	UniversityName      string    `json:"university_name"`
	UniversityLocation  string    `json:"university_location"`
	ProjectFieldOfStudy string    `json:"project_field_of_study"`
	ProjectMode         string    `json:"project_mode"`
	ProjectDuration     string    `json:"project_duration"`
	ProjectDetails      string    `json:"project_details"`
	Eligibility         string    `json:"eligibility"`
	Stipend             string    `json:"stipend"`
	ApplicationDeadline time.Time `json:"application_deadline"`
}
type Application struct {
    gorm.Model
    Name                  string    `json:"name"`
    RollNo                string    `json:"roll_no"`
    EmailId               string    `json:"email_id"`
    Programme             string    `json:"programme"`
    Department            string    `json:"department"`
    GraduationYear        string    `json:"graduation_year"`
    CPI                   float64   `json:"cpi"`
		ProjectID             uint      `json:"project_id"`
    ProjectName           string    `json:"project_name"`
    ProfessorName         string    `json:"professor_name"`
    ProfessorEmailId      string    `json:"professor_email_id"`
    UniversityName        string    `json:"university_name"`
    ApplicationDeadline   time.Time `json:"application_deadline"`
}