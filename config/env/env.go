package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func Get(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("error loading .env file")
	}

	return os.Getenv(key)
}