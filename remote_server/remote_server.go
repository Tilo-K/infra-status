package remote_server

import (
	"bytes"

	"github.com/charmbracelet/log"
	"golang.org/x/crypto/ssh"
)

func ExecuteCommand(conn *ssh.Client, command string) (stdout string, stderr string, err error) {
	session, err := conn.NewSession()
	if err != nil {
		log.Error("Could not create session", "Error", err)
		return "", "", err
	}
	defer session.Close()
	var stdoutBuf bytes.Buffer
	var stderrBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Stderr = &stderrBuf

	err = session.Run(command)
	if err != nil {
		log.Error("Could not execute command", "Command", command, "Error", err)
		return "", "", err
	}

	stdout = stdoutBuf.String()
	stderr = stderrBuf.String()
	return stdout, stderr, err
}

func DetermineOS(conn *ssh.Client) (string, error) {
	stdout, stderr, err := ExecuteCommand(conn, "uname -s")
	if stderr != "" {
		return "Windows", err
	}
	return stdout, err
}

func GetUptimeForServer(username string, address string) {
	conn, err := ConnectWithKey(username, address)
	if err != nil {
		log.Error("Could not connect to server", "Username", username, "Address", address, "Error", err)
	}

	hostSystem, err := DetermineOS(conn)
	if err != nil {
		log.Error("Could not determine OS", "Error", err)
		return
	}

	command := "uptime"
	if hostSystem == "Windows" {
		command = "systeminfo | findstr /i \"System Boot Time\""
	}

	stdout, stderr, err := ExecuteCommand(conn, command)
	if err != nil {
		log.Error("Could not get uptime", "Error", err)
		return
	}

	log.Info("Uptime", "Stdout", stdout, "Stderr", stderr)
}
