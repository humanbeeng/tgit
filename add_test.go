package main

import (
	"testing"
)

func TestSuccessfulAdd(t *testing.T) {
	files := []string{"asfasd.go"}
	err := add(files)
	if err != nil {
		t.Errorf("Unable to add main.go")
	}
}
