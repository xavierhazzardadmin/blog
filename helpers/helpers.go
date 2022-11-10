package helpers

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func SplitRows(s string) []string {
	return strings.Split(s, "\n\n")
}

func Get(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// set env variable using os package

	// return the env variable using os package
	return os.Getenv(key)
}

func GetEnv(key string) string {
    return os.Getenv(key)
}
