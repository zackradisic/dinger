package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Config represents the configuration details loaded from the config file
type Config struct {
	Sound string `json:"sound"`
}

// Configuration is the object that stores the Config for the user
var Configuration *Config

func defaultConfig() *Config {
	return &Config{
		Sound: os.Getenv("GOPATH") + "/src/github.com/zackradisic/dinger/sounds/ding.mp3",
	}
}

// ReadConfig reads the values of config.json and puts them into the Configuration variable,
// if not found it will create the file
func ReadConfig() {
	filename := os.Getenv("GOPATH") + "/src/github.com/zackradisic/dinger/config.json"
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		config := defaultConfig()
		jsonData, _ := json.Marshal(config)
		err = ioutil.WriteFile(filename, jsonData, 0644)
		if err != nil {
			log.Fatal("Error writing to config to file: " + err.Error())
		}
		Configuration = config
		return
	}
	config := &Config{}
	err = json.Unmarshal(bytes, config)
	if err != nil {
		log.Fatal(err)
	}

	Configuration = config
}
