package requests

import (
	"fmt"
	"testing"
)

func TestWhenConstructingSyncQueryThenTheQueryStringIsCorrectlyFormed(t *testing.T) {
	resourceTypes := ResourceTypes{"all"}
	syncQuery := NewQuery("token", "sync_token", resourceTypes)

	expectedString := fmt.Sprintf(`token=token&sync_token=sync_token&resource_types=%s`, resourceTypes.ToString())
	actualString := syncQuery.ToQueryString()

	if expectedString != actualString {
		t.Errorf("Expected '%s', but received '%s'", expectedString, actualString)
	}
}
