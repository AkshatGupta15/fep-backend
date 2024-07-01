package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pclubiitk/fep-backend/auth"
	"github.com/pclubiitk/fep-backend/mail"
	"github.com/pclubiitk/fep-backend/middleware"
	"github.com/spf13/viper"
)

func authServer(mail_channel chan mail.Mail) *http.Server {
	PORT := viper.GetString("PORT.AUTH")
	r := gin.New()
	r.Use(middleware.CORS())
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	auth.Router(mail_channel, r)
	
	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      r,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return server
}
