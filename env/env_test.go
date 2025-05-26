package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Config struct {
	AppName string `envconfig:"APP_NAME" required:"true"`
	Port    int    `envconfig:"PORT" default:"8080"`
}

func TestMust(t *testing.T) {
	t.Run("loads environment variables successfully", func(t *testing.T) {
		_ = os.Setenv("ENV", "local")
		_ = os.Setenv("APP_NAME", "TestApp")
		_ = os.Setenv("PORT", "8080")

		defer func() { _ = os.Unsetenv("ENV") }()
		defer func() { _ = os.Unsetenv("APP_NAME") }()
		defer func() { _ = os.Unsetenv("PORT") }()

		var config Config

		assert.NotPanics(t, func() {
			Must(&config)
		})

		assert.Equal(t, "TestApp", config.AppName)
		assert.Equal(t, 8080, config.Port)
	})

	t.Run("panics when required environment variable is missing", func(t *testing.T) {
		_ = os.Setenv("ENV", "local")

		defer func() { _ = os.Unsetenv("ENV") }()

		var config Config
		assert.Panics(t, func() {
			Must(&config)
		})
	})
}
