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

func main() {
	var dbversion int
	configPath, _ := lib.ReadEnvironmentVariable("CONFIG_PATH", lib.DefaultConfigPath)

	lib.ParseConfig(configPath, config)
	fmt.Println(config)

	dbversion, _ = lib.ChackDBVersion(config.DatabasePath)
	fmt.Println(dbversion)

	engine := lib.CreateGINEnvironment()

	// engine.GET("/:shortLink", GinRouteHandler)
	engine.Run(fmt.Sprintf(":%d", config.HTTPBindPort))
}
