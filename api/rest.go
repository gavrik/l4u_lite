package main

import (
	"fmt"
	"lib"

	"github.com/gin-gonic/gin"
)

// AuthTokenKey -
const AuthTokenKey = "AdminToken"

// REST - rest
type REST interface {
	Post(c *gin.Context)
	Get(c *gin.Context)
	Delete(c *gin.Context)
}

// NewLink - REST implementation for link section
func NewLink() REST {
	rest := new(LinkImplementation)
	rest.Db = lib.OpenDB(config.DatabasePath)
	return rest
}

// RESTErrorFunc -
func RESTErrorFunc(errNo int, errMsg string) RESTError {
	return RESTError{errNo, errMsg}
}

// GetAuthorizationToken -
func GetAuthorizationToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader != "" {
		return authHeader[6:len(authHeader)]
	}
	return ""
}

// TokenAuthorization -
func TokenAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHash := GetAuthorizationToken(c)
		fmt.Println(tokenHash)
		if token, ok := TokenCache[tokenHash]; ok {
			c.Set(AuthTokenKey, token)
		} else {
			c.Set(AuthTokenKey, AdminToken{})
		}
	}
}
