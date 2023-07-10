package main

import (
	"log"
	"runtime"

	"API/initializer"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvironment()
	log.Println(initializer.DBUser)
}

func main() {
	pc, _, _, _ := runtime.Caller(0)
	methodName := runtime.FuncForPC(pc).Name()
	log.Println(methodName, " Started")

	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})
	r.Run()

	log.Println(methodName, " Ended")
}
