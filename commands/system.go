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

// SystemClean esegue il comando sudo nix-collect-garbage -d.
func SystemClean() error {
	cmd := exec.Command("sh", "-c", "sudo nix-collect-garbage -d")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute system clean: %w", err)
	}

	return nil
}

// SystemTest esegue il comando nh os test.
func SystemTest() error {
	cmd := exec.Command("sh", "-c", "nh os test")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute system test: %w", err)
	}

	return nil
}

// SystemBuild esegue il comando nh os build.
func SystemBuild() error {
	cmd := exec.Command("sh", "-c", "nh os build")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute system build: %w", err)
	}

	return nil
}

// SystemBoot esegue il comando nh os boot.
func SystemBoot() error {
	cmd := exec.Command("sh", "-c", "nh os boot")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute system boot: %w", err)
	}

	return nil
}
