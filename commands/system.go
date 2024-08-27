package commands

import (
	"fmt"
	"os"
	"os/exec"
)

// CommandRunner is a type that represents a function that runs commands.
type CommandRunner func(name string, arg ...string) *exec.Cmd

// commandRunner is the actual function that runs commands.
// It can be overridden in tests.
var commandRunner CommandRunner = exec.Command

// executeCommand runs a system command with optional arguments.
func executeCommand(baseCmd string, args ...string) error {
	cmdStr := baseCmd

	// Check if `--impure` is present in the arguments
	for _, arg := range args {
		if arg == "--impure" {
			cmdStr += " -- --impure"
			break
		}
	}

	cmd := commandRunner("sh", "-c", cmdStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute command '%s': %w", cmdStr, err)
	}

	return nil
}

// SystemSwitch runs the command nh os switch.
func SystemSwitch(args ...string) error {
	return executeCommand("nh os switch", args...)
}

// SystemTest runs the command nh os test.
func SystemTest(args ...string) error {
	return executeCommand("nh os test", args...)
}

// SystemBuild runs the command nh os build.
func SystemBuild(args ...string) error {
	return executeCommand("nh os build", args...)
}

// SystemBoot runs the command nh os boot.
func SystemBoot(args ...string) error {
	return executeCommand("nh os boot", args...)
}

// SystemClean runs the command sudo nix-collect-garbage -d without additional options.
func SystemClean() error {
	// In this case, `--impure` is not applicable, so we directly use the command.
	cmd := commandRunner("sh", "-c", "sudo nix-collect-garbage -d")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute system clean: %w", err)
	}

	return nil
}
