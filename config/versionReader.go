package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	Versions Config
)

type Config struct {
	Data map[string]VersionConfig `json:"data"`
}

type VersionConfig struct {
	Platform     string          `json:"platform"`
	V            string          `json:"v"`
	Url          string          `json:"url"`
	UpdateConfig map[string]bool `json:"update"`
}

func ReadAll() bool {
	file, err := os.Open("./res/version.json")
	if err != nil {
		fmt.Println(err)
		return false
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Versions)
	if err != nil {
		fmt.Println(err)
	}

	return true
}
