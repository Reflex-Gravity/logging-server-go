package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	Logger *log.Logger
)

func main() {
	r := gin.Default()
	r.GET("/logger", func(ctx *gin.Context) {
		writeLog()
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()

}

func writeLog() {
	var filePath = "error.log"
	var file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println("Error", err)
	}
	Logger = log.New(file, "error:", log.Ldate|log.Ltime|log.Lshortfile)
	Logger.Println("asd")
}
