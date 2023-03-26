package main

import "time"

type ListItem struct {
	done        bool
	description string
}

type TodoList struct {
	items []ListItem
}

type DatedTodoList struct {
	list TodoList
	date time.Time
}

type ProjectTodoList struct {
	projectName string
	list        TodoList
}
