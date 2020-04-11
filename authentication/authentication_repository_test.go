package authentication

import (
	"testing"

	"github.com/kpdowns/todoist-cli/mocks"
)

func TestWhenRetrievingAccessTokenAndTheTokenExistsInStorageThenTheAccessTokenIsReturned(t *testing.T) {
	mockFile := &mocks.MockFile{
		Contents: "access-token",
	}

	repository := NewAuthenticationRepository(mockFile)

	_, err := repository.GetAccessToken()
	if err != nil {
		t.Errorf("Expected no error, but received '%s'", err.Error())
	}
}

func TestWhenRetrievingAccessTokenAndThereAreMultiplesLinesInStorageThenErrorIsReturned(t *testing.T) {
	mockFile := &mocks.MockFile{
		Contents: "access-token\ninvalid",
	}

	repository := NewAuthenticationRepository(mockFile)

	_, err := repository.GetAccessToken()
	if err == nil {
		t.Errorf("Expected '%s', but received '%s'", errorMalformedAuthenticationFile, err.Error())
	}
}

func TestWhenDeletingTheAccessTokenThenTheStorageIsCleared(t *testing.T) {
	mockFile := &mocks.MockFile{
		Contents: "access-token",
	}

	repository := NewAuthenticationRepository(mockFile)
	err := repository.DeleteAccessToken()
	if err != nil {
		t.Errorf("Expected no error, but received '%s'", err.Error())
	}

	if mockFile.Contents != "" {
		t.Errorf("Expected contents of the storage to be cleared")
	}
}

func TestWhenOverwritingTheAccessTokenThenTheContentsOfStorageIsTheNewAccessToken(t *testing.T) {
	mockFile := &mocks.MockFile{
		Contents: "access-token",
	}

	expectedContents := "test-token"
	repository := NewAuthenticationRepository(mockFile)
	err := repository.UpdateAccessToken(expectedContents)
	if err != nil {
		t.Errorf("Expected no error, but received '%s'", err.Error())
	}

	if mockFile.Contents != expectedContents {
		t.Errorf("Expected contents of the storage to be cleared but got '%s'", mockFile.Contents)
	}
}
