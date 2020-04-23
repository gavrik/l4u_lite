package main

import (
	"lib"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthTokenKey -
const AuthTokenKey = "AdminToken"

// REST - Gin Interface for REST calls
type REST interface {
	Post(c *gin.Context)
	Get(c *gin.Context)
	Delete(c *gin.Context)
	Patch(c *gin.Context)
}

// NewLink - REST implementation for link section
func NewLink() REST {
	rest := new(LinkImplementation)
	rest.Db = lib.OpenDB(config.DatabasePath)
	return rest
}

// RESTErrorFunc - Return error type
func RESTErrorFunc(errNo int, errMsg string) RESTError {
	return RESTError{errNo, errMsg}
}

// GetAuthorizationToken - Get token value from Authorization header
func GetAuthorizationToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader[:5] != "TOKEN" {
		return ""
	}
	if authHeader != "" {
		return authHeader[6:]
	}
	return ""
}

// TokenAuthorization - GIN middleware for API autorization
func TokenAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHash := GetAuthorizationToken(c)
		if token, ok := TokenCache[tokenHash]; ok {
			c.Set(AuthTokenKey, token)
		} else {
			c.Set(AuthTokenKey, AdminToken{})
		}
	}
}

// IsAuthorized - is rest call authorized for making request
func IsAuthorized(c *gin.Context) bool {
	token := c.MustGet(AuthTokenKey).(AdminToken)
	if token.Token == "" {
		c.JSON(http.StatusUnauthorized, RESTErrorFunc(1, "NotAuthorized"))
		return false
	}
	return true
}
