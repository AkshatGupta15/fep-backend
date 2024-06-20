package main

import (
	"log"
	"time"

	"github.com/pclubiitk/fep-backend/mail"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

const (
	readTimeout  = 5 * time.Second
	writeTimeout = 10 * time.Second
)

func main() {
	var g errgroup.Group
	mail_channel := make(chan mail.Mail)

	gin.SetMode(gin.ReleaseMode)

	go mail.Service(mail_channel)

	g.Go(func() error {
		return authServer(mail_channel).ListenAndServe()
	})
	g.Go(func() error {
		return studentServer(mail_channel).ListenAndServe()
	})

	log.Println("Server Started...")
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
