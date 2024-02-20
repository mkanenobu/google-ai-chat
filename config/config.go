package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	ApiKey string `json:"apiKey"`
}

func configLocation() string {
	return os.Getenv("HOME") + "/.config/google-ai-chat/config.json"
}

func NewConfig() Config {
	location := configLocation()

	file, err := os.ReadFile(location)
	if err != nil {
		panic(err)
	}

	var buf Config

	err = json.Unmarshal(file, &buf)
	if err != nil {
		panic(err)
	}

	return buf
}
