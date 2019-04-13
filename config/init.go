package config

import (
	"blog-api/app/utils"
	"gopkg.in/yaml.v2"
	"log"
	"sync"
)

const CONFIG_FILE_PATH = "./app.yaml"

var config_instance Config
var once_config_lock sync.Once

func GetConfig(path string) *Config {
	once_config_lock.Do(func() {
		config_byte := getConfigByte(path)
		err := yaml.Unmarshal([]byte(config_byte), &config_instance)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	})
	return &config_instance
}

func GetConfigNoSingleInstance(path string) *Config {
	config_byte := getConfigByte(path)
	err := yaml.Unmarshal([]byte(config_byte), &config_instance)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return &config_instance
}

func getConfigByte(path string) []byte {
	var config_byte []byte
	if path != "" {
		config_byte = utils.ReadFile(path)
	} else {
		config_byte = utils.ReadFile(CONFIG_FILE_PATH)
	}
	return config_byte
}
