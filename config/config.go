package config

import (
	"encoding/json"
	"os"
	"strings"
)

type RemoteServerConfig struct {
	Username string `json:"username"`
	Host     string `json:"host"`
}

type Config struct {
	RemoteServer []RemoteServerConfig `json:"remoteServer"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func WriteDefaultConfig(path string) error {
	parts := strings.Split(path, "/")
	dir := strings.Join(parts[:len(parts)-1], "/")

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	config := Config{
		RemoteServer: []RemoteServerConfig{
			{Username: "user1", Host: "host1"},
			{Username: "user2", Host: "host2"},
		},
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
