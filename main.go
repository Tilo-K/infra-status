package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"

	"tilok.dev/infra-status/config"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		_ = fmt.Errorf("Could not find home dir")
		homeDir = "/"
	}

	configPath := flag.String("config", filepath.Join(homeDir, ".config", "infra-status", "config.json"), "The path to your config file")
	flag.Parse()

	if _, err := os.Stat(*configPath); errors.Is(err, os.ErrNotExist) {
		config.WriteDefaultConfig(*configPath)
	}

	config, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Error("Error loading config", "Error", err)
		return
	}

	fmt.Println(config)
}
