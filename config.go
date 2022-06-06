package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbHost         string `json:"dbHost"`
	DbPort         string `json:"dbPort"`
	DbUser         string `json:"dbUser"`
	DbPassword     string `json:"dbPassword"`
	DbName         string `json:"dbName"`
	BotToken       string `json:"botToken"`
	DefGuildStatus bool   `json:"defaultGuildStatus"`
}

func LoadConfiguration(fileName string) *Config {
	var config Config
	configFile, err := os.Open(fileName)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return &config
}
