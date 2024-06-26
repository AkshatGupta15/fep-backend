package student

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	RollNo                 string  `gorm:"uniqueIndex" json:"roll_no"`
	Name                   string  `json:"name"`
	IITKEmail              string  `gorm:"uniqueIndex" json:"iitk_email"`
	CurrentCPI             float64 `json:"current_cpi"`
	UGCPI                  float64 `json:"ug_cpi"`
	ProgramDepartmentID    uint    `gorm:"index" json:"program_department_id"`
	ExpectedGraduationYear uint    `json:"expected_graduation_year"`
	Gender                 string  `json:"gender"`
	HasPassport            bool    `json:"has_passport"`
	IsEditable             bool    `json:"is_editable" gorm:"default:true"`
	IsVerified             bool    `json:"is_verified" gorm:"default:false"`
}
