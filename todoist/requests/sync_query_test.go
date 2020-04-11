package requests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSyncQuerySerialization(t *testing.T) {
	t.Run("Given a sync query, when converting the query to a query string, the string is a valid query string", func(t *testing.T) {
		resourceTypes := ResourceTypes{"all"}
		syncQuery := NewQuery("token", "sync_token", resourceTypes)

		expected := fmt.Sprintf(`token=token&sync_token=sync_token&resource_types=%s`, resourceTypes.ToString())
		actual := syncQuery.ToQueryString()

		assert.Equal(t, expected, actual)
	})
}
