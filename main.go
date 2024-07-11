package main

import (
	"fmt"
	"log"
	"os"

	"github.com/matteocavestri/cave/commands"
)

func main() {
	if len(os.Args) < 2 {
		printHelp("")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "--help":
		printHelp("")
	case "update":
		if len(os.Args) > 2 && os.Args[2] == "--help" {
			fmt.Println("Usage: cave update")
			os.Exit(0)
		}
		err := commands.Update()
		if err != nil {
			log.Fatalf("Update failed: %v", err)
		}
	case "system":
		if len(os.Args) < 3 {
			printHelp("system")
			os.Exit(1)
		}
		subcommand := os.Args[2]
		if subcommand == "--help" {
			printHelp("system")
			os.Exit(0)
		}
		switch subcommand {
		case "switch":
			err := commands.SystemSwitch()
			if err != nil {
				log.Fatalf("System switch failed: %v", err)
			}
		case "clean":
			err := commands.SystemClean()
			if err != nil {
				log.Fatalf("System clean failed: %v", err)
			}
		case "test":
			err := commands.SystemTest()
			if err != nil {
				log.Fatalf("System test failed: %v", err)
			}
		case "build":
			err := commands.SystemBuild()
			if err != nil {
				log.Fatalf("System build failed: %v", err)
			}
		case "boot":
			err := commands.SystemBoot()
			if err != nil {
				log.Fatalf("System boot failed: %v", err)
			}
		default:
			fmt.Printf("Unknown system subcommand: %s\n", subcommand)
			os.Exit(1)
		}
	case "home":
		if len(os.Args) < 3 {
			printHelp("home")
			os.Exit(1)
		}
		subcommand := os.Args[2]
		if subcommand == "--help" {
			printHelp("home")
			os.Exit(0)
		}
		switch subcommand {
		case "switch":
			err := commands.HomeSwitch()
			if err != nil {
				log.Fatalf("Home switch failed: %v", err)
			}
		case "build":
			err := commands.HomeBuild()
			if err != nil {
				log.Fatalf("Home build failed: %v", err)
			}
		default:
			fmt.Printf("Unknown home subcommand: %s\n", subcommand)
			os.Exit(1)
		}
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}

func printHelp(command string) {
	switch command {
	case "":
		fmt.Println("Usage: cave <command>")
		fmt.Println("Commands:")
		fmt.Println("  update      Update the flake")
		fmt.Println("  system      System commands")
		fmt.Println("  home        Home commands")
		fmt.Println("Use 'cave <command> --help' for more information about a command.")
	case "system":
		fmt.Println("Usage: cave system <subcommand>")
		fmt.Println("Subcommands:")
		fmt.Println("  switch      Switch the system configuration")
		fmt.Println("  clean       Clean the nix store")
		fmt.Println("  test        Test the system configuration")
		fmt.Println("  build       Build the system configuration")
		fmt.Println("  boot        Boot the system configuration")
	case "home":
		fmt.Println("Usage: cave home <subcommand>")
		fmt.Println("Subcommands:")
		fmt.Println("  switch      Switch the home configuration")
		fmt.Println("  build       Build the home configuration")
	default:
		fmt.Println("Unknown help topic.")
	}
}
