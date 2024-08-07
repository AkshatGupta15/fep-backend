package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pclubiitk/fep-backend/middleware"
	"github.com/pclubiitk/fep-backend/prof"
	"github.com/pclubiitk/fep-backend/project"
	"github.com/spf13/viper"
)

func adminProjectServer() *http.Server {
	PORT := viper.GetString("PORT.ADMIN.APPLICATION")
	engine := gin.New()
	engine.Use(middleware.CORS())
	engine.Use(middleware.Authenticator())
	engine.Use(middleware.EnsurePsuedoAdmin())
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

func adminProfServer() *http.Server {
	PORT := viper.GetString("PORT.ADMIN.PROF")
	engine := gin.New()
	engine.Use(middleware.CORS())
	engine.Use(middleware.Authenticator())
	engine.Use(middleware.EnsureAdmin())
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())

	prof.AdminRouter(engine)

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      engine,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return server
}
