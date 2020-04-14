package main

import (
	"lib"

	"github.com/gin-gonic/gin"
)

// DomainImplementation - implementation of link section
type DomainImplementation struct {
	Db lib.SQLiteDB
}

// Post -
func (impl *DomainImplementation) Post(c *gin.Context) {
	return
}

// Get -
func (impl *DomainImplementation) Get(c *gin.Context) {
	return
}

// Delete -
func (impl *DomainImplementation) Delete(c *gin.Context) {
	return
}

// Patch -
func (impl *DomainImplementation) Patch(c *gin.Context) {
	return
}
