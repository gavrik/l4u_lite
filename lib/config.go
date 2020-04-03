package lib

import (
	"errors"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

// ErrNoConfigFile -Error
var ErrNoConfigFile = errors.New("Config File doesn not exist")

// ParseConfig - yaml config parser
func ParseConfig(configFile string, configStruct interface{}) error {
	congNotExist := func(filename string) bool {
		_, err := os.Stat(filename)
		return os.IsNotExist(err)
	}
	if congNotExist(configFile) {
		return ErrNoConfigFile
	}
	b, err := ioutil.ReadFile(configFile)
	err = yaml.Unmarshal(b, configStruct)
	return err
}
