package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type shell int

const (
	shellUnknown shell = iota
	shellBash
	shellZsh
)

func detectShell() shell {
	s := os.Getenv("SHELL")
	if strings.Contains(s, "zsh") {
		return shellZsh
	}
	if strings.Contains(s, "bash") {
		return shellBash
	}
	return shellUnknown
}

func getHistoryPath(s shell) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not determine home directory: %w", err)
	}

	switch s {
	case shellZsh:
		return filepath.Join(home, ".zsh_history"), nil
	case shellBash:
		return filepath.Join(home, ".bash_history"), nil
	default:
		return "", fmt.Errorf("unsupported shell, please use bash or zsh")
	}
}

func GetPreviousCommand() (string, error) {
	s := detectShell()
	histPath, err := getHistoryPath(s)
	if err != nil {
		return "", err
	}

	file, err := os.Open(histPath)
	if err != nil {
		return "", fmt.Errorf("could not open history file: %w", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading history: %w", err)
	}

	if len(lines) == 0 {
		return "", fmt.Errorf("no commands found in history")
	}

	// Find the most recent command that isn't "aid"
	for i := len(lines) - 1; i >= 0; i-- {
		cmd := parseHistoryLine(lines[i], s)
		if !strings.HasPrefix(cmd, "aid") {
			return cmd, nil
		}
	}

	return "", fmt.Errorf("no commands found in history")
}

func parseHistoryLine(line string, s shell) string {
	// Zsh history format: ": timestamp:0;command"
	if s == shellZsh && strings.HasPrefix(line, ":") {
		parts := strings.SplitN(line, ";", 2)
		if len(parts) == 2 {
			return parts[1]
		}
	}
	return line
}
