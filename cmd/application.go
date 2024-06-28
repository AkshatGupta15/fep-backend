 package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/pclubiitk/fep-backend/middleware"
// 	"github.com/pclubiitk/fep-backend/application"
// 	"github.com/spf13/viper"
// )

// func applicationServer() *http.Server {
// 	PORT := viper.GetString("PORT.APPLICATION")
// 	engine := gin.New()
// 	engine.Use(middleware.CORS())
// 	engine.Use(gin.Recovery())
// 	engine.Use(gin.Logger())

// 	application.StudentRouter( engine)

// 	server := &http.Server{
// 		Addr:         ":" + PORT,
// 		Handler:      engine,
// 		ReadTimeout:  readTimeout,
// 		WriteTimeout: writeTimeout,
// 	}

// 	return server
// }