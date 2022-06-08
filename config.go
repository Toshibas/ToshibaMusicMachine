package main

import (
	
	"fmt"
	"os"

	"github.com/go-yaml/yaml" 

)

type Config struct {
	DbHost         string `yaml:"dbHost"`
	DbPort         string `yaml:"dbPort"`
	DbUser         string `yaml:"dbUser"`
	DbPassword     string `yaml:"dbPassword"`
	DbName         string `yaml:"dbName"`
	BotToken       string `yaml:"botToken"`
	DefGuildStatus bool   `yaml:"defaultGuildStatus"`
}

func LoadConfiguration(fileName string) *Config {
	var config Config
	configFile, err := os.Open(fileName)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	yamlParser := yaml.NewDecoder(configFile)
	yamlParser.Decode(&config)
	return &config
}
