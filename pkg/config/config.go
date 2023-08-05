package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const CONFIG_PATH = "./config.yaml"

type Config struct {
	Database struct {
		Host         string `yaml:"host"`
		Port         string `yaml:"port"`
		User         string `yaml:"user"`
		Password     string `yaml:"pass"`
		DatabaseName string `yaml:"dbname"`
	} `yaml:"database"`
	Listen string `yaml:"listen"`
}

func ReadConfig() Config {
	var config Config

	// Open YAML file
	file, err := os.Open(CONFIG_PATH)
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()

	// Decode YAML file to struct
	if file != nil {
		decoder := yaml.NewDecoder(file)
		if err := decoder.Decode(&config); err != nil {
			log.Println(err.Error())
			panic(err)
		}
	}

	return config
}

func (c Config) ComposeDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s",
		c.Database.Host, c.Database.Port, c.Database.User, c.Database.Password, c.Database.DatabaseName)
}
