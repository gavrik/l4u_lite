package main

import (
	"fmt"
	"lib"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GinRouteHandler - Default route handler. Return 301 with redirected link
func GinRouteHandler(c *gin.Context) {
	shortLink := c.Param("shortLink")
	var redirectLink string = ""

	if shortLink == "" {
		c.Redirect(http.StatusMovedPermanently, config.DefaultLink)
		return
	}

	db := lib.OpenDB(config.DatabasePath)
	defer db.Close()
	link := new(lib.LongLink)

	err := db.GetLongLink(shortLink, "", link)
	if err != nil {
		fmt.Println(err)
	}

	if link.LongLink != "" {
		redirectLink = link.LongLink
	} else {
		redirectLink = config.DefaultLink
	}

	c.Redirect(http.StatusMovedPermanently, redirectLink)
}
