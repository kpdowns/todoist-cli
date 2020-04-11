package requests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResourceTypeSerialization(t *testing.T) {
	t.Run("Given a resource type, when converting that resource type to a string, the resulting string is surrounded by double quotes", func(t *testing.T) {
		expected := `"all"`

		var resourceType ResourceType = "all"
		actual := resourceType.ToString()

		assert.Equal(t, expected, actual)
	})

	t.Run("Given a list of resource types, when converting those resources to a string, the resulting string is a json array", func(t *testing.T) {
		var resourceTypes ResourceTypes
		resourceTypes = append(resourceTypes, "projects")
		resourceTypes = append(resourceTypes, "items")

		expected := `["projects","items"]`
		actual := resourceTypes.ToString()

		if expected != actual {
			t.Errorf("Expected %s but got %s", expected, actual)
		}
	})
}
