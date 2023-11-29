package main

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"io/fs"
	"os"
)

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
