package project

import (
	"time"

	"github.com/pclubiitk/fep-backend/prof"
	"github.com/pclubiitk/fep-backend/student"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	ProjectName         string    `json:"project_name"`
	ProfID              uint      `json:"prof_id" gorm:"index;->;<-:create"`
	Prof                prof.Prof `json:"-" gorm:"foreignkey:ProfID"`
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
	ProjectID uint            `json:"project_id" gorm:"index;->;<-:create"`
	Project   Project         `json:"-" gorm:"foreignkey:ProjectID"`
	StudentID uint            `json:"student_id" gorm:"index;->;<-:create"`
	Student   student.Student `json:"-" gorm:"foreignkey:StudentID"`
}
