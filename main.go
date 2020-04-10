package main

import (
	"crypto/aes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Key string `json:"key"`
}

func main() {
	file, e := os.OpenFile("aes.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if e != nil {
		log.Fatalln("failed")
	}
	log.SetOutput(file)
	//key := "myverystrongpasswordo32bitlength"
	config, _ := LoadConfiguration("config.json")
	key1 := (config.Key)
	//fmt.Println(key1)
	log.Println(key1)
	fmt.Print("Enter 16 character text :")
	var plainText string
	fmt.Scan(&plainText)
	//fmt.Println(plainText)
	log.Println(plainText)
	ct := Encrypt([]byte(key1), plainText)
	log.Printf("Original Text:  %s\n", plainText)
	log.Printf("AES Encrypted Text:  %s\n", ct)
	log.Println(ct)
	dt := Decrypt([]byte(key1), ct)
	fmt.Printf("AES Decrypted Text:  %s\n", dt)
	log.Println(dt)

}
func LoadConfiguration(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}
func Encrypt(key []byte, plaintext string) string {
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Errorf("NewCipher(%d bytes) = %s", len(key), err)
		panic(err)
	}
	out := make([]byte, len(plaintext))
	c.Encrypt(out, []byte(plaintext))

	return hex.EncodeToString(out)
}

func Decrypt(key []byte, ct string) string {
	ciphertext, _ := hex.DecodeString(ct)
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Errorf("NewCipher(%d bytes) = %s", len(key), err)
		panic(err)
	}
	plain := make([]byte, len(ciphertext))
	c.Decrypt(plain, ciphertext)
	s := string(plain[:])
	return s
}
