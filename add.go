package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"io/fs"
	"os"
)

type FileType string

const (
	Blob FileType = "blob"
	Tree FileType = "tree"
)

type Staged struct {
	Filename string
	Hash     string
	FileType FileType
}

func add(fargs []string) error {
	idxinfo, err := os.Stat(".tgit/INDEX")
	if err != nil {
		return err
	}

	idxf, err := os.OpenFile(".tgit/INDEX", os.O_RDWR, fs.ModePerm)
	if err != nil {
		return err
	}
	defer idxf.Close()

	filesize := idxinfo.Size()

	staged := make(map[string]Staged)

	// Read INDEX file
	if filesize > 0 {
		b := make([]byte, filesize)

		_, err = idxf.Read(b)
		if err != nil {
			fmt.Printf("Unable to read %v", err)
			return err
		}

		buf := bytes.NewBuffer(b)

		dec := gob.NewDecoder(buf)

		if err := dec.Decode(&staged); err != nil {
			return fmt.Errorf("Unable to decode INDEX file, %v", err)
		}

	}

	// Stage
	for _, farg := range fargs {
		hash, err := fileSha1(farg)
		if err != nil {
			return err
		}

		if _, err := os.Stat(".tgit/objects/" + hash); errors.Is(err, fs.ErrExist) {
			fmt.Printf("No latest change in %v\n", farg)
			continue
		}

		if canStage(farg, hash, staged) {
			fmt.Printf("Added %v\n", farg)
			staged[farg] = Staged{Filename: farg, Hash: hash, FileType: Blob}
		}
	}

	// Write to INDEX
	buf := new(bytes.Buffer)
	g := gob.NewEncoder(buf)

	if err := g.Encode(staged); err != nil {
		return fmt.Errorf("Unable to perform staging")
	}

	if _, err := idxf.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("Unable to write to INDEX file")
	}

	return nil
}

func canStage(file string, fileHash string, staged map[string]Staged) bool {
	if _, err := os.Stat(".tgit/objects/" + fileHash); err == nil {
		fmt.Printf("No latest change in %v\n", file)
		return false
	}

	if _, ok := staged[file]; ok {
		fmt.Printf("%v already added\n", file)
		return false
	}

	return true
}
