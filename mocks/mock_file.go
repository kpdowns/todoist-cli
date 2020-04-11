package mocks

// MockFile is a structure containing a representation of the contents of storage in a string
type MockFile struct {
	Contents string
}

// ReadContents returns the contents the in-memory string containing the file contents
func (f *MockFile) ReadContents() (string, error) {
	return f.Contents, nil
}

// OverwriteContents overwrites all contents of the in-memory string
func (f *MockFile) OverwriteContents(contents string) error {
	f.Contents = contents
	return nil
}
