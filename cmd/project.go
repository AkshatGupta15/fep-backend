package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pclubiitk/fep-backend/middleware"
	"github.com/pclubiitk/fep-backend/project"
	"github.com/spf13/viper"
)

func adminProjectServer() *http.Server {
	PORT := viper.GetString("PORT.PROJECT")
	engine := gin.New()
	engine.Use(middleware.CORS())
	engine.Use(middleware.Authenticator())
	engine.Use(middleware.EnsureAdmin())
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())

	project.AdminRouter(engine)

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      engine,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return server
}
