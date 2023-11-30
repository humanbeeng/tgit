package main

import (
	"testing"
)

func TestSuccessfulCommit(t *testing.T) {
	msg := "Test commit"
	err := commit(msg)
	if err != nil {
		t.Errorf("Commit failed")
	}
}
