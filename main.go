package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	server := gin.New()

	APP_PORT := os.Getenv("APP_PORT")
	fmt.Println("---- APP ----")
	fmt.Println("APP_PORT:", APP_PORT)
	fmt.Println("")

	if err := server.Run(":" + os.Getenv("APP_PORT")); err != nil {
		logrus.WithError(err).Fatal("Error starting server")
	}
}
