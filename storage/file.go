package storage

import (
	"bufio"
	"fmt"
	"os"
)

const (
	errorFailedToAccessFile = "Failed to access file located at '%s'"
)

// File is a facade in front of raw filesystem access. Files are opened in read/write.
type File interface {
	ReadContents() (string, error)
	OverwriteContents(contents string) error
}

type file struct {
	path string
}

// NewFile creates a new facade
func NewFile(path string) File {
	return &file{
		path: path,
	}
}

// ReadContents returns the contents of the file as a string
func (f *file) ReadContents() (string, error) {
	file, err := f.openFile()
	if err != nil {
		return "", fmt.Errorf(errorFailedToAccessFile, f.path)
	}

	defer file.Close()

	var contents string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents = contents + scanner.Text()
	}

	return contents, nil
}

// OverwriteContents overwrites all contents of the file
func (f *file) OverwriteContents(contents string) error {
	file, err := f.openFile()
	if err != nil {
		return fmt.Errorf(errorFailedToAccessFile, f.path)
	}

	defer file.Close()

	err = file.Truncate(0)
	if err != nil {
		return err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	_, err = file.WriteString(contents)
	return err
}

func (f *file) openFile() (*os.File, error) {
	file, err := os.OpenFile(f.path, os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		return nil, err
	}
	return file, nil
}
