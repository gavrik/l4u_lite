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
	IsDebug      bool   `yaml:"isDebug"`
	HTTPBindPort int    `yaml:"httpBindPort"`
	DatabasePath string `yaml:"databasePath"`
	DefaultLink  string `yaml:"defaultLink"`
	APIHost      string `yaml:"apiHost"`
	IsCreateDB   bool   `yaml:"isCreateDB"`
}

// TokenCache - Admin Tokens Cache
var TokenCache map[string]AdminToken

func populateAdminCache(dbPath string) {
	TokenCache = make(map[string]AdminToken)
	db := NewAPIDB(dbPath)
	defer db.Close()

	err := db.GetAdminTokens(TokenCache)
	if err != nil {
		panic(err)
	}
}

func main() {
	var dbversion int
	var err error
	configPath, _ := lib.ReadEnvironmentVariable("CONFIG_PATH", lib.DefaultConfigPath)

	lib.ParseConfig(configPath, config)
	if config.IsDebug {
		fmt.Println(config)
	}

	if !lib.IsFileExixts(config.DatabasePath) && config.IsCreateDB {
		fmt.Println("Create fresh database structure")
		err = CreateFreshDB(config)
		if err != nil {
			panic(err)
		}
	}

	dbversion, err = lib.ChackDBVersion(config.DatabasePath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Database structure version: %d\n", dbversion)

	populateAdminCache(config.DatabasePath)
	if len(TokenCache) == 0 {
		panic("Admin Token Cache are empty. Exiting.")
	}
	if config.IsDebug {
		fmt.Println(TokenCache)
	}

	engine := lib.CreateGINEnvironment(config.IsDebug)
	engine.Use(TokenAuthorization())
	linkRoutes := engine.Group("/link")
	//adminRoutes := engine.Group("/admin")
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
	linkRoutes.GET("/info/:domain/:link_hash", link.Get)
	// Get all link info depending to domain
	//engine.GET("/link/info/:domain", link.Get)
	// Create new short link
	linkRoutes.POST("/create", link.Post)
	// Delete short link
	linkRoutes.DELETE("/delete/:domain/:link_hash", link.Delete)
	// Update link parameters
	linkRoutes.PATCH("/patch", link.Patch)

	//
	engine.Run(fmt.Sprintf(":%d", config.HTTPBindPort))
}
