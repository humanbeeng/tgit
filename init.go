package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func initRepo() error {
	// create .tgit folder
	err := os.Mkdir(".tgit", os.ModePerm)
	if err != nil {
		if errors.Is(err, fs.ErrExist) {
			return fmt.Errorf("tgit repository already exists")
		}
		return fmt.Errorf("Unable to init repository %v", err)
	}

	contents := []string{"objects/", "refs/", "refs/heads/"}
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

	if _, err := os.Create(".tgit/refs/HEAD"); err != nil {
		return fmt.Errorf("Unable to create HEAD file")
	}

	if _, err := os.Create(".tgit/refs/heads/main"); err != nil {
		return fmt.Errorf("Unable to create main branch file")
	}

	if err := os.WriteFile(".tgit/refs/HEAD", []byte("main"), fs.ModePerm); err != nil {
		return err
	}

	fmt.Println("Initialised an empty tgit repository")
	return nil
}
