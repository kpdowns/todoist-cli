package requests

import "testing"

func TestStringRepresentationOfResourceTypeContainsQuotationMarks(t *testing.T) {
	var resourceType ResourceType = "all"

	expectedStringRepresentation := `"all"`
	actualStringRepresentation := resourceType.ToString()

	if expectedStringRepresentation != actualStringRepresentation {
		t.Errorf("Expected %s but got %s", expectedStringRepresentation, actualStringRepresentation)
	}
}

func TestGivenAListOfResourceTypesWhenConvertingToStringThenEachElementIsCommandSeparatedWithTheWholeStringSurroundedByBrackets(t *testing.T) {
	var resourceTypes ResourceTypes
	resourceTypes = append(resourceTypes, "projects")
	resourceTypes = append(resourceTypes, "items")

	expectedStringRepresentation := `["projects","items"]`
	actualStringRepresentation := resourceTypes.ToString()

	if expectedStringRepresentation != actualStringRepresentation {
		t.Errorf("Expected %s but got %s", expectedStringRepresentation, actualStringRepresentation)
	}
}
