package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func fileSha1(file string) (string, error) {
	contents, err := os.ReadFile(file)
	if err != nil {
		return "", fmt.Errorf("Unable to read %s %d", file, err)
	}

	hash := getSha1(contents)
	return hash, nil
}

func getSha1(data []byte) string {
	hasher := sha1.New()
	hasher.Write(data)

	// Get the final hash and convert it to a hexadecimal string
	hashInBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashInBytes)
	return hashString
}

func repoExists() bool {
	_, err := os.Stat(".tgit")

	if errors.Is(err, fs.ErrNotExist) {
		if err != nil {
			return false
		}
	}
	return true
}

func currBranch(headInf fs.FileInfo) (string, error) {
	headFile, err := os.OpenFile(".tgit/refs/HEAD", os.O_RDWR, fs.ModePerm)
	if err != nil {
		return "", fmt.Errorf("Unable to read HEAD file")
	}
	defer headFile.Close()

	// Get the curr branch name
	var currBranch string
	if headInf.Size() > 0 {
		sc := bufio.NewScanner(headFile)
		for sc.Scan() {
			currBranch = sc.Text()
		}
	}
	return currBranch, nil
}
