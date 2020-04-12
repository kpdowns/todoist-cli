package mocks

// MockFile is a structure containing a representation of the contents of storage in a string
type MockFile struct {
	Contents       string
	ReadError      error
	OverwriteError error
}

// ReadContents returns the contents the in-memory string containing the file contents
func (f *MockFile) ReadContents() (string, error) {
	if f.ReadError != nil {
		return "", f.ReadError
	}
	return f.Contents, nil
}

// OverwriteContents overwrites all contents of the in-memory string
func (f *MockFile) OverwriteContents(contents string) error {
	if f.OverwriteError != nil {
		return f.OverwriteError
	}
	f.Contents = contents
	return nil
}
