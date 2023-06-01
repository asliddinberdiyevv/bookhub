package main

import (
	"bookhub/db"
	"flag"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	server := gin.New()
	flag.Parse()
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.WithField("port", os.Getenv("PORT")).Debug("starting server")

	DB := db.ConnectDB()
	defer DB.Close()

	RS := db.ConnectRS()
	defer RS.Close()

	server.GET("/", func(ctx *gin.Context) {
		logrus.Info("get request done")
		ctx.JSON(http.StatusOK, gin.H{
			"status": "successful",
		})
	})

	if err := server.Run(":" + os.Getenv("PORT")); err != nil {
		logrus.WithError(err).Fatal("Error starting server")
	}
}
