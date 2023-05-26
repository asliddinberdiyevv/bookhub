package main

import (
	"bookhub/db"
	"flag"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	server := gin.New()
	flag.Parse()
	logrus.SetLevel(logrus.DebugLevel)
	logrus.WithField("port", os.Getenv("PORT")).Debug("starting server")

	DB := db.ConnectDB()
	defer DB.Close()

	RS := db.ConnectRS()
	defer RS.Close()

	if err := server.Run(":" + os.Getenv("PORT")); err != nil {
		logrus.WithError(err).Fatal("Error starting server")
	}
}
