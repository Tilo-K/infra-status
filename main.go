package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		_ = fmt.Errorf("Could not find home dir")
		homeDir = "/"
	}

	configPath := flag.String("config", filepath.Join(homeDir, ".config", "infra-status", "config.json"), "The path to your config file")
	flag.Parse()

	fmt.Println(*configPath)
}
