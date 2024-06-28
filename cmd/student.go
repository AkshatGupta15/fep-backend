package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pclubiitk/fep-backend/application"
	"github.com/pclubiitk/fep-backend/mail"
	"github.com/pclubiitk/fep-backend/middleware"
	"github.com/pclubiitk/fep-backend/student"
	"github.com/spf13/viper"
)

func studentServer(mail_channel chan mail.Mail) *http.Server {
	PORT := viper.GetString("PORT.STUDENT")
	engine := gin.New()
	engine.Use(middleware.CORS())
	engine.Use(middleware.Authenticator())
	engine.Use(gin.Logger())
	student.StudentRouter(engine)
	// rc.StudentRouter(engine)
	application.StudentRouter(engine)

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      engine,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return server
}
