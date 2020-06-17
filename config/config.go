package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Config represents the configuration details loaded from the config file
type Config struct {
	Sound string `json:"sound"`
}

// Configuration is the object that stores the Config for the user
var (
	Configuration *Config
	filename      = os.Getenv("GOPATH") + "/src/github.com/zackradisic/dinger/config.json"
)

func defaultConfig() *Config {
	return &Config{
		Sound: os.Getenv("GOPATH") + "/src/github.com/zackradisic/dinger/sounds/ding.mp3",
	}
}

// SetValue sets a config's value for the given key
func SetValue(key string, val interface{}) error {
	switch key {
	case "sound":
		if s, ok := val.(string); ok {
			err := setSound(s)
			if err != nil {
				return fmt.Errorf("Error setting sound: %s", err.Error())
			}
		} else {
			return fmt.Errorf("Sound must be a valid path")
		}

	}

	WriteConfig()
	return nil
}

// PrintValue prints the value of the config option to stdout at the given key if it exists
func PrintValue(key string) error {
	var input map[string]interface{}
	out, _ := json.Marshal(Configuration)
	json.Unmarshal(out, &input)

	if val, ok := input[key]; ok && val != nil {
		fmt.Printf("Currently, (%s) is set to: (%s)\n", key, val)
		return nil
	}

	return fmt.Errorf("config option (%s) does not exist", key)
}

func setSound(path string) error {
	_, err := os.Open(path)
	if err != nil {
		return err
	}

	Configuration.Sound = path
	return nil
}

// WriteConfig writes the values currently stored in the Configuration variable to config.json
func WriteConfig() {
	jsonData, _ := json.Marshal(Configuration)
	err := ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		log.Fatal("Error writing config to file: " + err.Error())
	}
}

// ReadConfig reads the values of config.json and puts them into the Configuration variable,
// if not found it will create the file
func ReadConfig() {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		config := defaultConfig()
		jsonData, _ := json.Marshal(config)
		err = ioutil.WriteFile(filename, jsonData, 0644)
		if err != nil {
			log.Fatal("Error writing config to file: " + err.Error())
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
