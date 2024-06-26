package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pclubiitk/fep-backend/mail"
	"github.com/sirupsen/logrus"
)

type signUpRequest struct {
	UserID   string `json:"user_id" binding:"required"` // roll or PF number
	Password string `json:"password" binding:"required"`
	UserOTP  string `json:"user_otp" binding:"required"`
}

func signUpHandler(mail_channel chan mail.Mail) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var signupReq signUpRequest

		if err := ctx.ShouldBindJSON(&signupReq); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		verified, err := verifyOTP(ctx, signupReq.UserID, signupReq.UserOTP)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if !verified {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Roll No OTP"})
			return
		}

		hashedPwd := hashAndSalt(signupReq.Password)

		id, err := firstOrCreateUser(ctx, &User{
			UserID:   signupReq.UserID,
			Password: hashedPwd,
		})

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		logrus.Infof("User %s created successfully with id %d", signupReq.UserID, id)
		mail_channel <- mail.GenerateMail(signupReq.UserID, "Registered on HC Automation Portal", "Dear "+signupReq.UserID+",\n\nYou have been registered on HC Automation Portal")
		ctx.JSON(http.StatusOK, gin.H{"status": "Successfully signed up"})
	}
}
