package main

import (
	"errors"
	"fmt"
	"lib"
)

const appName string = "l4uApplication"
const defaultConfigPath string = "./config.yaml"
const databaseVersion int = 1

var config = new(AppConfig)

// ErrWrongDBVersion - Error
var ErrWrongDBVersion = errors.New("Wrong database version")

// AppConfig - Application config
type AppConfig struct {
	Version      int    `yaml:"version"`
	Kind         string `yaml:"kind"`
	HTTPBindPort int    `yaml:"httpBindPort"`
	DatabasePath string `yaml:"databasePath"`
	DefaultLink  string `yaml:"defaultLink"`
}

func main() {
	var dbversion int
	configPath, err := lib.ReadEnvironmentVariable("CONFIG_PATH", defaultConfigPath)

	lib.ParseConfig(configPath, config)
	fmt.Println(config)

	db := OpenDB(config.DatabasePath)
	dbversion, err = db.CheckDBversion()
	if err != nil {
		panic(err)
	}
	if dbversion != databaseVersion {
		panic(ErrWrongDBVersion.Error())
	}
	db.Close()

	engine := lib.CreateGINEnvironment()

	engine.GET("/:shortLink", GinRouteHandler)
	engine.Run(fmt.Sprintf(":%d", config.HTTPBindPort))

}
