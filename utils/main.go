package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"net/url"
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

func GeneratePassword(length int) string {
	// Define the character set for your random string
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Create a slice to store the random bytes
	randomBytes := make([]byte, length)

	// Calculate the maximum index in the character set
	maxIndex := big.NewInt(int64(len(charSet)))

	for i := 0; i < length; i++ {
		// Generate a random index
		randomIndex, _ := rand.Int(rand.Reader, maxIndex)
		randomBytes[i] = charSet[randomIndex.Int64()]
	}

	return string(randomBytes)
}

func VerifyCaptcha(verificationToken string) (bool, error) {
	secret := os.Getenv("HCAPTCHA_SECRET")

	postData := url.Values{
		"secret":   {secret},
		"response": {verificationToken},
	}

	// Create a new POST request to the hCaptcha verify URL with the required data
	res, err := http.PostForm("https://hcaptcha.com/siteverify", postData)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	// Decode the JSON response from hCaptcha
	var resBody structs.HCaptchaResponse
	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		return false, err
	}

	return resBody.Success, nil
}

func Encrypt(plainText string) (string, error) {
	key := os.Getenv("ENCRYPTION_KEY")

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// Create a new GCM cipher mode instance
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(plainText), nil)

	return hex.EncodeToString(ciphertext), nil
}

func Decrypt(encrypted string) (string, error) {
	key := os.Getenv("ENCRYPTION_KEY")

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	encryptedData, err := hex.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(encryptedData) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := encryptedData[:nonceSize], encryptedData[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
