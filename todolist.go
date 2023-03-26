package main

// ListItem represents a single item in a TodoList
type ListItem struct {
	// done indicates whether an item is completed or not
	done bool

	// name is a short-name description of the item that is always shown in the list
	name string

	// description is an expanded description of the item, hidden by default
	description string

	// estimate is an estimated amount of effort (minutes, points, etc.) to complete the item.
	estimate int
}

type TodoList struct {
	// identifier is used to identify the TodoList for saving and loading
	identifier string

	// items contains the individual ListItems.
	items []ListItem
}
