package utils

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Thund3rD3v/SuperGuardian/structs"
)

func GetConfig() structs.Config {
	var config structs.Config

	configf, err := os.Open("config.json")
	if err != nil {
		log.Println("Error while loading config.json file," + err.Error())
		return config
	}
	defer configf.Close()

	decoder := json.NewDecoder(configf)
	err = decoder.Decode(&config)
	if err != nil {
		log.Println("Error while reading config.json file," + err.Error())
		return config
	}

	return config
}
