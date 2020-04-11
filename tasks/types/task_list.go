package types

import "sort"

// TaskList is a list of unordered tasks
type TaskList []Task

// SortByDueDateThenSortByPriority sorts the slice of tasks by due date, then priority. Returns a new slice of tasks.
func (t TaskList) SortByDueDateThenSortByPriority() TaskList {
	daysSeen := make(map[int64]int64)
	tasksGroupedByDay := make(map[int64]TaskList)
	for _, task := range t {
		day := task.DueDate.Unix()
		if _, dateAlreadySeen := daysSeen[day]; !dateAlreadySeen {
			daysSeen[day] = day
		}

		tasksGroupedByDay[day] = append(tasksGroupedByDay[day], task)
	}

	days := make([]int64, 0, len(daysSeen))
	for day := range daysSeen {
		days = append(days, day)
	}

	sort.Slice(days, func(i, j int) bool { return days[i] < days[j] })

	var tasksGroupedByDayAndOrdererByPriority TaskList
	for _, day := range days {
		sort.Sort(tasksGroupedByDay[day])
		tasksGroupedByDayAndOrdererByPriority = append(tasksGroupedByDayAndOrdererByPriority, tasksGroupedByDay[day]...)
	}

	return tasksGroupedByDayAndOrdererByPriority
}

// Len returns the length of the TaskList
func (t TaskList) Len() int { return len(t) }

// Less returns true if the priority of the task is higher than the one being compared
func (t TaskList) Less(i, j int) bool {
	return t[i].Priority > t[j].Priority
}

// Swap swaps two different tasks in the slice
func (t TaskList) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
