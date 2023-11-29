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
			err := initRepository()
			if err != nil {
				fmt.Println(err)
			}
		}
	case "help":
		{
			displayHelp()
		}

	case "add":
		{
			err := add(args[2:])
			if err != nil {
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

func initRepository() error {
	// create .tgit folder
	err := os.Mkdir(".tgit", os.ModePerm)
	if err != nil {
		return fmt.Errorf("Unable to init repository")
	}

	contents := []string{"objects/", "refs/"}
	_, err = os.Create(".tgit/INDEX")
	if err != nil {
		return err
	}

	for _, c := range contents {
		err := os.MkdirAll(".tgit/"+c, os.ModePerm)
		if err != nil {
			return fmt.Errorf("Unable to create %s", c)
		}
	}
	fmt.Println("Initialised an empty tgit repository")
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
