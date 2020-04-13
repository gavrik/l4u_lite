package main

import (
	"fmt"
	"lib"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LinkImplementation - implementation of link section
type LinkImplementation struct {
	Db lib.SQLiteDB
}

// Post - Create
// curl --data '{"shortLink": "test-link_1"}' --header "Authorization: TOKEN e319e2a5-95f5-48fd-bbc9-7315df21c382" http://127.0.0.1:8081/link/create
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
	fmt.Println(err)

	c.JSON(http.StatusOK, link)
}

// Get - Get info
func (impl *LinkImplementation) Get(c *gin.Context) {
	if !IsAuthorized(c) {
		return
	}
	var link Link
	shortLink := c.Param("link_hash")
	domain := c.Param("domain")
	db := NewAPIDBro(config.DatabasePath)
	defer db.Close()
	link.Domain = domain
	err := db.GetLink(shortLink, &link)
	fmt.Println(err)

	c.JSON(http.StatusOK, link)
}

// Delete - delete
func (impl *LinkImplementation) Delete(c *gin.Context) {
	return
}
