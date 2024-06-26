package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pclubiitk/fep-backend/mail"
)

func Router(mail_channel chan mail.Mail, r *gin.Engine) {
	auth := r.Group("/api/auth")
	{
		auth.GET("/hello", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"msg": "hello"})
		})
		auth.POST("/otp", otpHandler(mail_channel))
		auth.POST("/signup", signUpHandler(mail_channel))
	}
}
