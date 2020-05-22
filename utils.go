package main

import (
	"io/ioutil"
	"os"
)

// FileBytes returns bytes of file.
func FileBytes(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return []byte{}, err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return []byte{}, err
	}

	return b, nil
}
