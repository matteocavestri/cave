package commands

import (
	"fmt"
	"os"
	"os/exec"
)

// Update esegue il comando di aggiornamento del flake.
func Update() error {
	flakeDir := os.Getenv("FLAKE")
	if flakeDir == "" {
		return fmt.Errorf("FLAKE environment variable is not set")
	}

	// Esegui i comandi tramite una shell
	shellCommands := fmt.Sprintf("cd %s && nix flake update && cd -", flakeDir)

	cmd := exec.Command("sh", "-c", shellCommands)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute commands: %w", err)
	}

	return nil
}
