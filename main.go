package main

import (
	"flag"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	server := gin.New()
	flag.Parse()
	logrus.SetLevel(logrus.DebugLevel)
	logrus.WithField("port", os.Getenv("APP_PORT")).Debug("starting server")

	if err := server.Run(":" + os.Getenv("APP_PORT")); err != nil {
		logrus.WithError(err).Fatal("Error starting server")
	}
}
