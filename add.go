package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"io/fs"
	"os"
)

type Staged struct {
	Filename string
	Hash     string
	// filetype ?
}

func add(fargs []string) error {
	if !repoExists() {
		return fmt.Errorf("Repository not initialized")
	}

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
		fmt.Println("File size > 0")
		b := make([]byte, filesize)

		_, err = idxf.Read(b)
		if err != nil {
			fmt.Printf("Unable to read %v", err)
			return err
		}

		buf := bytes.NewBuffer(b)

		dec := gob.NewDecoder(buf)

		if err := dec.Decode(&staged); err != nil {
			return fmt.Errorf("Unable to decode INDEX file")
		}

		fmt.Printf("Read from index file %v", staged)
	}

	// Stage
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

		if _, ok := staged[hash]; ok {
			fmt.Printf("%v already staged\n", farg)
		} else {
			fmt.Printf("Added %v\n", farg)
			staged[hash] = Staged{Filename: farg, Hash: hash}
		}
	}

	// Write to INDEX
	buf := new(bytes.Buffer)
	g := gob.NewEncoder(buf)

	if err := g.Encode(staged); err != nil {
		return fmt.Errorf("Unable to perform staging")
	}

	if err := os.Truncate(".tgit/INDEX", 0); err != nil {
		return fmt.Errorf("Unable to clear contents of INDEX file %v", err)
	}

	if _, err := idxf.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("Unable to write to INDEX file")
	}

	fmt.Println("Staging completed")

	return nil
}
