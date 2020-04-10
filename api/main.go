package main

import (
	"fmt"
	"lib"
)

// AppName -
const AppName string = "l4uAPIAplication"

var config = new(AppConfig)

// AppConfig - Application config
type AppConfig struct {
	Version      int    `yaml:"version"`
	Kind         string `yaml:"kind"`
	HTTPBindPort int    `yaml:"httpBindPort"`
	DatabasePath string `yaml:"databasePath"`
	DefaultLink  string `yaml:"defaultLink"`
	APIHost      string `yaml:"apiHost"`
}

// TokenCache - Admin Tokens Cache
var TokenCache map[string]AdminToken

func populateAdminCache(dbPath string) {
	TokenCache = make(map[string]AdminToken)
	db := NewAPIDB(dbPath)
	defer db.Close()

	err := db.GetAdminTokens(TokenCache)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(TokenCache)

}

func main() {
	var dbversion int
	configPath, _ := lib.ReadEnvironmentVariable("CONFIG_PATH", lib.DefaultConfigPath)

	lib.ParseConfig(configPath, config)
	fmt.Println(config)

	dbversion, _ = lib.ChackDBVersion(config.DatabasePath)
	fmt.Println(dbversion)

	populateAdminCache(config.DatabasePath)

	engine := lib.CreateGINEnvironment()
	linkRoutes := engine.Group("/link", TokenAuthorization())
	link := NewLink()

	// Create new token
	//engine.POST("/admin/create")
	// Delete token. Root token could not be deleted.
	//engine.DELETE("/admin/delete/:hash_token")
	// Get token info
	//engine.GET("/admin/info")
	// Create new domain record
	//engine.POST("/domain/create")
	// Delete domain record with all links and tokens depended on it
	//engine.DELETE("/domain/delete/:domain")
	// Get domain info
	//engine.GET("/domain/info/:domain")
	// Get link info
	//engine.GET("/link/info/:domain/:link_hash")
	// Get all link info depending to domain
	//engine.GET("/link/info/:domain")
	// Create new short link
	linkRoutes.POST("/create", link.Post)
	// Delete short link
	//engine.DELETE("/link/delete/:domain/:link_hash")

	//
	engine.Run(fmt.Sprintf(":%d", config.HTTPBindPort))
}
