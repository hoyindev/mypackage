package ioutil

import (
	// "errors"
	// "fmt"
	// "sync"
	"testing"
)

func TestReadDirNames(t *testing.T) {
	nameDir := "./"
	s, err := ReadDirNames(nameDir, true)
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	t.Error(s)
}

func TestIsDirEmpty(t *testing.T) {
	nameDir := "/"
	s, err := IsDirEmpty(nameDir)
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	t.Error(s)
}
