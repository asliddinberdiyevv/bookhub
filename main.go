package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	server := gin.New()
	flag.Parse()
	logrus.SetLevel(logrus.DebugLevel)
	logrus.WithField("port", os.Getenv("APP_PORT")).Debug("starting server")

	server.GET("/", func(ctx *gin.Context) {
		logrus.Info("get request")
		ctx.JSON(http.StatusOK, gin.H{"message": "main server"})
	})

	if err := server.Run(os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT")); err != nil {
		logrus.WithError(err).Fatal("Error starting server")
	}
}
