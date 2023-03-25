package main

type ListItem struct {
	done        bool
	description string
}

type TodoList struct {
	items []ListItem
}

// type DatedTodoList struct {
// 	list TodoList
// 	date Time
// }

type ProjectTodoList struct {
	projectName string
	list        TodoList
}
