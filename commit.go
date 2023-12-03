package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"io/fs"
	"os"
)

/*
Git commit flow
Check the branch head if it has any head

if head exists, add that to the tree-object
*/

type CommitObject struct {
	Message     string
	TreeHash    string
	SubtreeHash string
}

func commit(args []string) error {
	// Check if message is passed
	if len(args) == 0 {
		return fmt.Errorf("No commit message found. Aborting!")
	}

	// Get staged files
	idx, err := os.ReadFile(".tgit/INDEX")
	if err != nil {
		return err
	}

	if len(idx) == 0 {
		return fmt.Errorf("Nothing to commit")
	}

	// Get the current branch
	headInf, err := os.Stat(".tgit/refs/HEAD")
	if err != nil {
		return err
	}

	currBranch, err := currBranch(headInf)
	if err != nil {
		return err
	}

	if _, err := os.Stat(".tgit/refs/heads/" + currBranch); err != nil {
		return err
	}

	cbFile, err := os.OpenFile(".tgit/refs/heads/"+currBranch, os.O_RDWR, fs.ModePerm)
	if err != nil {
		return fmt.Errorf("Unable to read curr branch file")
	}

	defer cbFile.Close()

	// Get the latest commit information from curr branch
	var latestHash string
	if headInf.Size() > 0 {
		sc := bufio.NewScanner(cbFile)
		for sc.Scan() {
			latestHash = sc.Text()
		}
	}

	var staged map[string]Staged

	idxbuf := bytes.NewBuffer(idx)
	idxdec := gob.NewDecoder(idxbuf)

	if err := idxdec.Decode(&staged); err != nil {
		return fmt.Errorf("Unable to decode INDEX file, %v", err)
	}

	stageHash := getSha1(idx)

	// Create a commit struct
	cmtMsg := args[0]
	co := CommitObject{
		Message:     cmtMsg,
		TreeHash:    stageHash,
		SubtreeHash: latestHash,
	}

	fmt.Println("About to commit", co)

	// Create hash for the commit struct
	cmtHash := getSha1([]byte(cmtMsg + stageHash + latestHash))

	for stagedFile := range staged {
		hash, err := fileSha1(stagedFile)
		if err != nil {
			return err
		}
		f, err := os.ReadFile(stagedFile)
		if err != nil {
			return err
		}

		if err := os.WriteFile(".tgit/objects/"+hash, f, fs.ModePerm); err != nil {
			fmt.Println("Aborting commit")
			return err
		}

	}

	// Write commit-object to a file
	cmtFile, err := os.OpenFile(".tgit/objects/"+cmtHash, os.O_CREATE|os.O_RDWR, fs.ModePerm)
	if err != nil {
		return fmt.Errorf("Unable to open commit object, %v", err)
	}
	defer cmtFile.Close()

	cobuf := new(bytes.Buffer)
	cog := gob.NewEncoder(cobuf)

	if err := cog.Encode(co); err != nil {
		return err
	}

	if _, err := cmtFile.Write(cobuf.Bytes()); err != nil {
		return fmt.Errorf("Unable to commit, %v", err)
	}

	// Reset staging area
	if err = os.Truncate(".tgit/INDEX", 0); err != nil {
		return err
	}

	// Update the head to the latest commit hash
	if err := os.WriteFile(".tgit/refs/heads/"+currBranch, []byte(cmtHash), fs.ModePerm); err != nil {
		return fmt.Errorf("Unable to update latest commit to branch")
	}

	fmt.Println("Committed")
	return nil
}
