package utils

import (
	"crypto/rand"
	"encoding/json"
	"log"
	"math/big"
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

func WriteConfig(config structs.Config) error {
	jsonConfig, err := json.Marshal(config)
	if err != nil {
		return err
	}

	err = os.WriteFile("config.json", jsonConfig, 0644)
	if err != nil {
		return err
	}

	return nil
}

func GeneratePassword(l int) string {
	// Define the character set for your random string
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Create a slice to store the random bytes
	randomBytes := make([]byte, l)

	// Calculate the maximum index in the character set
	maxIndex := big.NewInt(int64(len(charSet)))

	for i := 0; i < l; i++ {
		// Generate a random index
		randomIndex, _ := rand.Int(rand.Reader, maxIndex)
		randomBytes[i] = charSet[randomIndex.Int64()]
	}

	return string(randomBytes)
}
