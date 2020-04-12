package types

import "testing"

func TestGivenATaskWhenConvertingToStringThenThePriorityIsAStringCorrespondingToTheValue(t *testing.T) {
	var tasksToTest = []struct {
		task           Task
		expectedString string
	}{
		{
			Task{
				ID:       1,
				Priority: 1,
				Content:  "test",
			},
			"1\tLow\ttest",
		},
		{
			Task{
				ID:       2,
				Priority: 2,
				Content:  "test2",
			},
			"2\tNormal\ttest2",
		},
		{
			Task{
				ID:       3,
				Priority: 3,
				Content:  "test3",
			},
			"3\tUrgent\ttest3",
		},
		{
			Task{
				ID:       4,
				Priority: 4,
				Content:  "test4",
			},
			"4\tVery Urgent\ttest4",
		},
	}

	for _, taskToTest := range tasksToTest {
		stringRepresentation := taskToTest.task.AsString()
		if stringRepresentation != taskToTest.expectedString {
			t.Errorf("Expected '%s', got '%s'", taskToTest.expectedString, stringRepresentation)
		}
	}
}
