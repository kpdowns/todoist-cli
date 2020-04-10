package requests

import (
	"fmt"
)

// ResourceType is the type of resource work is being performed on when interacting with the Todoist API
type ResourceType string

// ResourceTypes is a slice of ResourceType
type ResourceTypes []ResourceType

// ToString converts the resource type into a representation that can be provided to the Todoist API
func (r *ResourceType) ToString() string {
	return fmt.Sprintf(`"%s"`, string(*r))
}

// ToString converts the slice to a comma separated list
func (r *ResourceTypes) ToString() string {
	stringRepresentation := "["
	for index, resourceType := range *r {
		stringRepresentation += resourceType.ToString()
		if index < len(*r)-1 {
			stringRepresentation += ","
		}
	}
	stringRepresentation += "]"
	return stringRepresentation
}
