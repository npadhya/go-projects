package initializer

import (
	"log"
	"os"
	"runtime"

	"github.com/joho/godotenv"
)

var DBUser string

func LoadEnvironment() {
	pc, _, _, _ := runtime.Caller(0)
	methodName := runtime.FuncForPC(pc).Name()
	log.Println(methodName, " Started")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DBUser = os.Getenv("DBUser")
	log.Println(methodName, " Ended")
}
