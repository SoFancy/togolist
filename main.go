package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. Load todolist
	fmt.Println("Welcome to togolist!")

	// Get command-line arguments
	args := os.Args[1:]
	if len(args) > 0 {
		fmt.Println("Args!")
	} else {
		fmt.Println("No argument specified, creating new todolist for today if none exists")
	}
}
