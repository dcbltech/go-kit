package env

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func Must[E any](e *E) {
	ee := strings.ToLower(os.Getenv("ENV"))

	if ee == "" || ee == "local" {
		_ = godotenv.Load(".env")
	}

	err := envconfig.Process("", e)
	if err != nil {
		panic(err)
	}
}
