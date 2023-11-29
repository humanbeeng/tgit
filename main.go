package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("No tgit command found")
		displayHelp()
		return
	}

	switch args[0] {
	case "init":
		{
			err := initRepository()
			if err != nil {
				fmt.Println(err)
			}
		}
	case "help":
		{
			displayHelp()
		}
	default:
		{
			fmt.Println("Invalid command", args[0])
			displayHelp()
		}
	}
}

func initRepository() error {
	// create .tgit folder
	err := os.Mkdir(".tgit", os.ModePerm)
	if err != nil {
		return fmt.Errorf("Unable to init repository")
	}

	contents := []string{"objects/", "refs/"}

	for _, c := range contents {
		err := os.MkdirAll(".tgit/"+c, os.ModePerm)
		if err != nil {
			return fmt.Errorf("Unable to create %s", c)
		}
	}
	return nil
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
