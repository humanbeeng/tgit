package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/fs"
	"os"
)

func checkout(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("No branch name found.")
	}

	if !branchExists(args[0]) {
		fmt.Println(args[0], "does not exists")
		return fmt.Errorf("Branch %v does not exists", args[0])
	}

	// Get the latest commit hash from the to checkout branch
	lch, err := os.ReadFile(".tgit/refs/heads/" + args[0])
	if err != nil {
		return fmt.Errorf("Unable to get latest commit hash, %v", err)
	}

	if len(lch) > 0 {

		q := make([]string, 0)
		q = append(q, string(lch))
		uniqueFiles := make(map[string]string, 0)

		for len(q) != 0 {
			currHash := q[0]
			q = q[1:]

			// Get the commit object
			cFile, err := os.ReadFile(".tgit/objects/" + string(currHash))
			if err != nil {
				return err
			}

			if len(cFile) == 0 {
				return fmt.Errorf("No commits found. Aborting")
			}

			var cobj CommitObject

			buf := bytes.NewBuffer(cFile)
			dec := gob.NewDecoder(buf)

			if err := dec.Decode(&cobj); err != nil {
				return fmt.Errorf("Unable to decode commit object , %v", err)
			}

			// Get staged tree obj which is a map[filename]TreeObj from the currHash
			cTree, err := os.ReadFile(".tgit/objects/" + cobj.TreeHash)
			if err != nil {
				return err
			}

			var tree map[string]TreeItem

			treeBuf := bytes.NewBuffer(cTree)
			treeDec := gob.NewDecoder(treeBuf)

			if err := treeDec.Decode(&tree); err != nil {
				return fmt.Errorf("Unable to decode tree object map, %v", err)
			}

			// Check if the filename already exists in map. If not, add to map
			for filename, t := range tree {
				if _, ok := uniqueFiles[filename]; !ok {
					uniqueFiles[filename] = t.Hash
				}
			}

			if cobj.SubtreeHash != "" {
				q = append(q, cobj.SubtreeHash)
			}

		}

		for filename, hash := range uniqueFiles {
			source, err := os.ReadFile(".tgit/objects/" + hash)
			if err != nil {
				return err
			}

			if err := os.WriteFile(filename, source, fs.ModePerm); err != nil {
				return err
			}

		}
	}

	// Reset HEAD
	if err = os.Truncate(".tgit/refs/HEAD", 0); err != nil {
		return err
	}
	// Update HEAD
	if err := os.WriteFile(".tgit/refs/HEAD", []byte(args[0]), fs.ModePerm); err != nil {
		return err
	}
	fmt.Println("Checked out to ", args[0])

	return nil
}
