package utils

import (
	"io"
	"os"
)

// Analyze receive a file name and a function which will process the content of that file
// Allows line by line reads
func Analyze(file string, f func(handle io.Reader) error) error {
	handle, err := os.Open(file)

	if err != nil {
		return err
	}
	defer handle.Close()
	return f(handle)
}
