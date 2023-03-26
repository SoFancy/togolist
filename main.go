package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func openTodolist(projectName string) {
	fmt.Println("Doing stuff for the project", projectName)
}

func main() {
	// 1. Load todolist
	fmt.Println("Welcome to togolist!")

	now := time.Now()

	// Get command-line arguments
	args := os.Args[1:]
	if len(args) > 0 {
		switch strings.ToLower(args[0]) {
		case "project":
			fmt.Println("Project")
			if len(args) < 2 {
				fmt.Println("No projectname specified. I don't know what to do. Panic attack. Aborting.")
				os.Exit(1)
			}
			openTodolist(args[1])
		case "date":
			fmt.Println("Date")
			if len(args) < 2 {
				fmt.Println("You wanted to open a togolist for a date, but no date was specified. Defaulting to the beginning of the universe. Error. Aborting.")
				os.Exit(1)
			}
			openTodolist(args[1])
		case "yesterday":
			fmt.Println("Yesterday")
			yesterday := now.AddDate(0, 0, -1).Format("2006-01-02")
			openTodolist(yesterday)
		case "today":
			fmt.Println("Today")
			today := now.Format("2006-01-02")
			openTodolist(today)
		default:
			fmt.Println("Unknown argument:", args[0])
		}
	} else {
		fmt.Println("No argument specified, creating new todolist for today if none exists")
		today := now.Format("2006-01-02")
		openTodolist(today)
	}
}
