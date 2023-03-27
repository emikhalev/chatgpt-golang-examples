package utils

import (
	"io/ioutil"
	"os"
	"sync"
)

var (
	apiKeyConfigFile = "../config/apiKey.txt"
	once             sync.Once

	apiKey = ""
)

func init() {
	if cfgFile := os.Getenv("chatgpt-golang-examples-apikey-config"); cfgFile != "" {
		apiKeyConfigFile = cfgFile
	}
}

func APIKey() string {
	readConfig()
	return apiKey
}

func readConfig() {
	once.Do(func() {
		if b, err := ioutil.ReadFile(apiKeyConfigFile); err == nil {
			apiKey = string(b)
		}
	})
}
