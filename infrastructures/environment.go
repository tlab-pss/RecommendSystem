package infrastructures

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitEnvironment() {
	InitEnvWithPath("")
}

func InitEnvWithPath(path string) {
	err := godotenv.Load(fmt.Sprintf(path+"%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}
