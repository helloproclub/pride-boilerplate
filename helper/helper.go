package helper

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(variable string, fallback interface{}) interface{} {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	res := os.Getenv(variable)
	if res != "" {
		return res
	}
	return fallback
}
