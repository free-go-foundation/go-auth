package config

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"io/ioutil"
)

type Config struct {
	Env     Environment   `json:"Environment"`
	DB      Database      `json:"Database"`
	Signing Authorization `json:"Signing"`
}

type Environment struct {
	Port string `json:"Port"`
	Host string `json:"Host"`
}

type Database struct {
	Host string `json:"Host"`
}

type Authorization struct {
	SecretKey string `json:"SecretKey"`
}

var DevConfig *Config = nil

func init() {
	absPath, _ := filepath.Abs("../go-auth/config/config.json")
	file, _ := ioutil.ReadFile(absPath)
	err := json.Unmarshal(file, &DevConfig)
	if err != nil {
		fmt.Println("error:", err)
	}
}
