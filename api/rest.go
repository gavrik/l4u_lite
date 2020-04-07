package main

import (
	"github.com/gin-gonic/gin"
)

// REST - rest
type REST interface{
	Post(c *gin.Context)
	Get(c *gin.Context)
	Delete(c *gin.Context)
}

// NewLink - REST implementation for link section
func NewLink() REST {
	rest := new(LinkImplementation)
	return rest
}
