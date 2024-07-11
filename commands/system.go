package commands

import (
	"fmt"
	"os"
	"os/exec"
)

// SystemSwitch esegue il comando nh os switch.
func SystemSwitch() error {
	cmd := exec.Command("sh", "-c", "nh os switch")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute system switch: %w", err)
	}

	return nil
}
