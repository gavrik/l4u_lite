package lib

import (
	"errors"
	"os"

	"github.com/gin-gonic/gin"
	uuid "github.com/gofrs/uuid"
)

// ErrEnvNotExist - Error
var ErrEnvNotExist = errors.New("Lib: Environment variable is not exist")

// ErrNoUUID - Error
var ErrNoUUID = errors.New("Can't generate UUID")

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

// GetUUID - Generate unique UUID
func GetUUID() (string, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return "", ErrNoUUID
	}
	return uuid.String(), nil
}
