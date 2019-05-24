package config

import (
	"encoding/json"
	"io/ioutil"
)

var cfg *Config

type Config struct {
	Addresses   []string `json:"addresses"`
	NodeURL     string   `json:"node_url"`
	BipPriceUSD float64  `json:"bip_price_usd"`
}

func initConfig() {
	configData, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(configData, &cfg); err != nil {
		panic(err)
	}
}

func Get() *Config {
	if cfg == nil {
		initConfig()
	}

	return cfg
}
