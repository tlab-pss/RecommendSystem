package infrastructures

import (
	"log"

	"github.com/joho/godotenv"
)

func InitEnvironment() {
	InitEnvWithPath("")
}

func InitEnvWithPath(path string) {
	err := godotenv.Load(path + ".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}
