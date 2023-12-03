package main

import (
	"fmt"
	"io/fs"
	"os"
)

func branch(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Branch name required\n")
	}

	// Get the current branch
	headInf, err := os.Stat(".tgit/refs/HEAD")
	if err != nil {
		return err
	}

	cb, err := currBranch(headInf)
	if err != nil {
		return err
	}

	cbFile, err := os.OpenFile(".tgit/refs/heads/"+cb, os.O_RDWR, fs.ModePerm)
	if err != nil {
		return fmt.Errorf("Unable to read curr branch file")
	}
	defer cbFile.Close()

	// Get the latest commit hash from the curr branch
	lch, err := os.ReadFile(".tgit/refs/heads/" + cb)
	if err != nil {
		return fmt.Errorf("Unable to get latest commit hash, %v", err)
	}

	// Create new branch
	if _, err := os.Create(".tgit/refs/heads/" + args[0]); err != nil {
		return fmt.Errorf("Unable to create main branch file")
	}

	// Write latest commit hash to new branch file
	if err := os.WriteFile(".tgit/refs/heads/"+args[0], lch, fs.ModePerm); err != nil {
		return fmt.Errorf("Unable to write latest commit hash")
	}

	// Reset HEAD
	if err = os.Truncate(".tgit/refs/HEAD", 0); err != nil {
		return err
	}

	// Update HEAD
	if err := os.WriteFile(".tgit/refs/HEAD", []byte(args[0]), fs.ModePerm); err != nil {
		return err
	}

	fmt.Printf("%v created \n", args[0])

	return nil
}
