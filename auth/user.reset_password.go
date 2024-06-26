package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pclubiitk/fep-backend/mail"
	"github.com/sirupsen/logrus"
)

type resetPasswordRequest struct {
	UserID      string `json:"user_id" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
	OTP         string `json:"otp" binding:"required"`
}

func resetPasswordHandler(mail_channel chan mail.Mail) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var resetPasswordReq resetPasswordRequest

		if err := ctx.ShouldBindJSON(&resetPasswordReq); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		verified, err := verifyOTP(ctx, resetPasswordReq.UserID, resetPasswordReq.OTP)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !verified {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid OTP"})
			return
		}

		hashedPwd := hashAndSalt(resetPasswordReq.NewPassword)

		ok, err := updatePassword(ctx, resetPasswordReq.UserID, hashedPwd)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !ok {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No such user exists"})
			return
		}

		logrus.Infof("Password of %s reset successfully", resetPasswordReq.UserID)
		mail_channel <- mail.GenerateMail(resetPasswordReq.UserID, "Password Reset Successful", "Your password has been reset successfully.")

		ctx.JSON(http.StatusOK, gin.H{"status": "Successfully reset password"})
	}
}
