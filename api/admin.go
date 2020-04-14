package main

import (
	"lib"

	"github.com/gin-gonic/gin"
)

// AdminImplementation - implementation of link section
type AdminImplementation struct {
	Db lib.SQLiteDB
}

// Post -
func (impl *AdminImplementation) Post(c *gin.Context) {
	return
}

// Get -
func (impl *AdminImplementation) Get(c *gin.Context) {
	return
}

// Delete -
func (impl *AdminImplementation) Delete(c *gin.Context) {
	return
}

// Patch -
func (impl *AdminImplementation) Patch(c *gin.Context) {
	return
}
