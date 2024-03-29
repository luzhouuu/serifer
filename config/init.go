package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Configuration definition
type Configuration struct {
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	Expand   string `json:"expand"`
	Database struct {
		Driver     string `json:"driver"`
		Connection string `json:"connection"`
	} `json:"database"`
	API struct {
		USE string `json:"use"`
	} `json:"api"`
}

var configuration Configuration

// Load loads Configuration
func Load() {
	// what's err
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(bytes, &configuration)
}

// Get return configuration
func Get() *Configuration {
	return &configuration
}
