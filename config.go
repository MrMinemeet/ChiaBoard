package main

import (
	"encoding/json"
	"log"
	"os"
)

type TConfig struct {
	ChiaPath string
}

func LoadSettings() TConfig {
	// Open file "TConfig.json" and read it into TConfig struct
	configJson, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	var config TConfig
	json.Unmarshal([]byte(configJson), &config)

	return config
}

const TmpChiaPath = "/usr/bin/chia"
