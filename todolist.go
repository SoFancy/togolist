package main

type ListItem struct {
	done        bool
	description string
}

type TodoList struct {
	identifier string
	items      []ListItem
}
