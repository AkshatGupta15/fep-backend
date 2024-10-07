package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pclubiitk/fep-backend/constants"
	"github.com/pclubiitk/fep-backend/mail"
	"github.com/pclubiitk/fep-backend/middleware"
	"github.com/pclubiitk/fep-backend/prof"
)

type profSignUpRequest struct {
	Password         string `json:"password"`
	ProfessorName    string `json:"professor_name"`
	ProfessorEmailId string `json:"professor_email_id"`
	UniversityName   string `json:"university_name"`
}

func profSignUpHandler(mail_channel chan mail.Mail) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		middleware.Authenticator()(ctx)
		if middleware.GetUserID(ctx) == "" {
			return
		}

		if middleware.GetRoleID(ctx) != constants.GOD {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Only God  can sign up for PROF"})
			return
		}

		var request profSignUpRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if request.ProfessorName == "" || request.ProfessorEmailId == "" || request.UniversityName == "" || request.Password == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ProfessorName, ProfessorEmailId, password and UniversityName are required"})
			return
		}
		newProf := prof.Prof{
			ProfessorName:    request.ProfessorName,
			ProfessorEmailId: request.ProfessorEmailId,
			UniversityName:   request.UniversityName,
		}
		err := prof.CreateProf(ctx, &newProf)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var UserReq User
		UserReq.UserID = request.ProfessorEmailId
		// password = GeneratePassword()
		password := request.Password
		UserReq.Password = hashAndSalt(password)
		UserReq.RoleID = constants.PROF
		UserReq.Name = request.ProfessorName

		id, err := firstOrCreateUser(ctx, &UserReq)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		mail_channel <- mail.GenerateMail(UserReq.UserID, "New Credentials generated", "Your new credentials are: \n\nUser ID: "+UserReq.UserID+"\nPassword: "+UserReq.Password+"\n\nYou can reset the password from // To be added")
		// <a href= \"https://anc.iitk.ac.in/reset-password\">here</a>
		ctx.JSON(http.StatusOK, gin.H{"id": id})
	}
}
