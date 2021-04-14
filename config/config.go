package config

import (
	"encoding/json"
	"io/ioutil"
)

type ServerConfig struct {
	ServerAddr string
	TimeoutSecond int
}

type Config struct {
	Server ServerConfig
}

func NewConfig(data []byte) (*Config, error) {
	cfg := &Config{}
	err := cfg.Unmarshal(data)
	
	return cfg, err
}

func NewConfigFromFile(file string) (*Config, error){
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return NewConfig(data)
}

func (cfg *Config) Marshal() ([]byte, error) {
	return json.Marshal(cfg)
}

func (cfg *Config) Unmarshal(data []byte) error {
	return json.Unmarshal(data, cfg)
}
