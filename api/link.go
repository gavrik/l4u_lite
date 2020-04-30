package main

import (
	"lib"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LinkImplementation - implementation of link section
type LinkImplementation struct {
	Db lib.SQLiteDB
}

// Post - Create new link
func (impl *LinkImplementation) Post(c *gin.Context) {
	if !IsAuthorized(c) {
		return
	}
	var link Link
	c.BindJSON(&link)
	if link.ShortLink == "" {
		link.ShortLink = lib.RandomString(lib.RandomStringHashLength)
	}
	db := NewAPIDB(config.DatabasePath)
	defer db.Close()
	err := db.PutLink(&link)
	if err != nil {
		c.JSON(http.StatusNotFound, RESTErrorFunc(5, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, link)
}

// Get - Get link info
func (impl *LinkImplementation) Get(c *gin.Context) {
	if !IsAuthorized(c) {
		return
	}
	var link Link
	link.ShortLink = c.Param("link_hash")
	link.Domain = c.Param("domain")
	db := NewAPIDBro(config.DatabasePath)
	defer db.Close()
	err := db.GetLink(&link)
	if err != nil {
		c.JSON(http.StatusNotFound, RESTErrorFunc(5, err.Error()))
		return
	}
	if link.LongLink == "" {
		c.JSON(http.StatusNotFound, RESTErrorFunc(6, "LinkInfoNotFound"))
	}

	c.JSON(http.StatusOK, link)
}

// Delete - Delete link from database
func (impl *LinkImplementation) Delete(c *gin.Context) {
	if !IsAuthorized(c) {
		return
	}
	var link Link
	link.ShortLink = c.Param("link_hash")
	link.Domain = c.Param("domain")
	db := NewAPIDB(config.DatabasePath)
	defer db.Close()
	err := db.DeleteLink(&link)
	if err != nil {
		c.JSON(http.StatusNotFound, RESTErrorFunc(5, err.Error()))
		return
	}

	c.JSON(http.StatusOK, RESTErrorFunc(2, "HashLinkDeleted"))
}

// Patch - Update link parameters
func (impl *LinkImplementation) Patch(c *gin.Context) {
	if !IsAuthorized(c) {
		return
	}
	var link Link
	c.BindJSON(&link)
	if link.ShortLink == "" {
		c.JSON(http.StatusBadRequest, RESTErrorFunc(4, "HashLinkBadRequest"))
		return
	}
	db := NewAPIDB(config.DatabasePath)
	defer db.Close()
	err := db.UpdateLink(&link)
	if err != nil {
		c.JSON(http.StatusNotFound, RESTErrorFunc(5, err.Error()))
		return
	}

	c.JSON(http.StatusOK, RESTErrorFunc(3, "HashLinkUpdated"))
}
