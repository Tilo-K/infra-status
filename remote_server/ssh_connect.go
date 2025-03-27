package remote_server

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/ssh"
)

// Function to attempt SSH connection with a given key file
func tryKey(user string, address string, keyPath string) (*ssh.Client, error) {
	// Read the private key from file
	key, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read private key %s: %w", keyPath, err)
	}

	// Create the signer for the private key
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, fmt.Errorf("unable to parse private key %s: %w", keyPath, err)
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", address, config)
	if err != nil {
		return nil, fmt.Errorf("unable to connect with key %s: %w", keyPath, err)
	}

	return conn, nil
}

func ConnectWithKey(username string, address string) (*ssh.Client, error) {
	if !strings.Contains(address, ":") {
		address += ":22"
	}

	usr, err := user.Current()
	if err != nil {
		fmt.Println("Unable to determine current user:", err)
		return nil, err
	}

	keyPaths := []string{
		filepath.Join(usr.HomeDir, ".ssh", "id_rsa"),
		filepath.Join(usr.HomeDir, ".ssh", "id_dsa"),
		filepath.Join(usr.HomeDir, ".ssh", "id_ecdsa"),
		filepath.Join(usr.HomeDir, ".ssh", "id_ed25519"),
	}

	for _, keyPath := range keyPaths {
		conn, err := tryKey(username, address, keyPath)
		if err == nil {
			return conn, nil
		}
	}

	return nil, errors.New("failed to connect with any of the provided keys")
}
