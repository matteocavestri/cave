package commands

import (
	"fmt"
	"os"
	"os/exec"
)

// HomeSwitch esegue il comando nh home switch.
func HomeSwitch() error {
	cmd := exec.Command("sh", "-c", "nh home switch")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute home switch: %w", err)
	}

	return nil
}

// HomeSwitch esegue il comando nh home build.
func HomeBuild() error {
	cmd := exec.Command("sh", "-c", "nh home build")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute home build: %w", err)
	}

	return nil
}
