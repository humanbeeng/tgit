package main

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"
)

func add(fargs []string) error {
	if !repoExists() {
		return fmt.Errorf("Repository not initialized")
	}

	idxf, err := os.OpenFile(".tgit/INDEX", os.O_RDWR, fs.ModeAppend)
	if err != nil {
		return err
	}
	defer idxf.Close()

	sc := bufio.NewScanner(idxf)
	stagedHashes := make(map[string]bool)

	for sc.Scan() {
		stagedHashes[sc.Text()] = true
	}

	// Check all files
	for _, farg := range fargs {
		f, err := os.Open(farg)
		if err != nil {
			return fmt.Errorf("Unable to open %s %d", farg, err)
		}
		defer f.Close()

		contents, err := io.ReadAll(f)
		if err != nil {
			return fmt.Errorf("Unable to read %s %d", farg, err)
		}

		hash := getSha1(contents)

		if _, ok := stagedHashes[hash]; ok {
			fmt.Printf("%v already staged\n", farg)
		} else {
			if _, err := idxf.WriteString(hash + "\n"); err != nil {
				fmt.Println(err)
				return fmt.Errorf("Unable to stage %v\n", farg)
			}
			fmt.Printf("%v added\n", farg)
		}
	}

	return nil
}
