package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// WriteMode is the default
// file/folder writing mode
const WriteMode = 0755

// DoesFileExist is a quick way to check if
// a file is already in the filesystem
func DoesFileExist(filename string) (bool, error) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

// GetFileBytes opens a file and gets all its contents
func GetFileBytes(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// IsFileNotFoundError determines whether the
// error is from the file not existing
func IsFileNotFoundError(err error) bool {
	str := err.Error()
	return strings.Contains(str, ": no such file or directory")
}

// WriteFileBytes is a convenience method for writing
// to a file
func WriteFileBytes(filename string, bytes []byte) error {
	return ioutil.WriteFile(filename, bytes, WriteMode)
}

// WriteBashScript is a quick way to write
// a bash script
func WriteBashScript(location string, content string) error {
	contents := fmt.Sprintf("#!/bin/sh\n%s", content)
	return ioutil.WriteFile(location, []byte(contents), 0755)
}
