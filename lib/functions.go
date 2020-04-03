package lib

import (
	"errors"
	"os"

	"github.com/gin-gonic/gin"
)

// ErrEnvNotExist - Error
var ErrEnvNotExist = errors.New("Lib: Environment variable is not exist")

// ReadEnvironmentVariable - read environment variable
func ReadEnvironmentVariable(key string, defaultValue string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue, ErrEnvNotExist
	}
	return val, nil
}

// CreateGINEnvironment - Create default GIN engine and return it
func CreateGINEnvironment() *gin.Engine {
	router := gin.Default()
	return router
}
