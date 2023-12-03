package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 1 {
		fmt.Println("No tgit command found")
		displayHelp()
		return
	}

	switch args[1] {
	case "init":
		{
			err := initRepo()
			if err != nil {
				fmt.Println(err)
			}
		}
	case "help":
		{
			if !repoExists() {
				fmt.Println("Repository not initialized")
				return
			}
			displayHelp()
		}

	case "checkout":
		{
			checkout(args[2:])
		}

	case "add":
		{
			if !repoExists() {
				fmt.Println("Repository not initialized")
				return
			}

			err := add(args[2:])
			if err != nil {
				fmt.Println(err)
			}
		}

	case "commit":
		{
			if !repoExists() {
				fmt.Println("Repository not initialized")
				return
			}

			if err := commit(args[2:]); err != nil {
				fmt.Println(err)
			}
		}

	case "branch":
		{
			if !repoExists() {
				fmt.Println("Repository not initialized")
				return
			}

			if err := branch(args[2:]); err != nil {
				fmt.Println(err)
			}
		}

	default:
		{
			fmt.Println("Invalid command", args[0])
			displayHelp()
		}
	}
}

func displayHelp() {
	fmt.Println("Tiny Git by humanbeeng!")
	fmt.Println("Available commands:")
	fmt.Println("\tinit: Initializes an empty tgit repository ")
	fmt.Println("\tadd: Adds files to staging")
	fmt.Println("\tcommit: Commit staged files")
	fmt.Println("\tbranch: Creates a new branch")
	fmt.Println("\tcheckout: Checkout to branch")
	fmt.Println("\tclone: Clones a remote repository")
	fmt.Println("\thelp: Display this help message")
}
