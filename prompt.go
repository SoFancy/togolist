package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

var options = []string{
	"Option 1",
	"Option 2",
	"Option 3",
	"Option 4",
	"Option 5",
}

func drawOptions(selectedIndex int) {
	for i, option := range options {
		selector := ' '
		if i == selectedIndex {
			selector = '>'
		}
		tbprint(0, i, selector, option)
	}
	termbox.Flush()
}

func tbprint(x, y int, prefix rune, msg string) {
	for _, ch := range msg {
		termbox.SetCell(x, y, ch, termbox.ColorDefault, termbox.ColorDefault)
		x++
	}
	termbox.SetCell(0, y, prefix, termbox.ColorDefault, termbox.ColorDefault)
}

func openPrompt(projectName string) {
	// Initialize termbox
	err := termbox.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer termbox.Close()

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	selectedIndex := 0
	drawOptions(selectedIndex)

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowUp:
				if selectedIndex > 0 {
					selectedIndex--
					drawOptions(selectedIndex)
				}
			case termbox.KeyArrowDown:
				if selectedIndex < len(options)-1 {
					selectedIndex++
					drawOptions(selectedIndex)
				}
			case termbox.KeyEnter:
				break mainloop
			case termbox.KeyEsc, termbox.KeyCtrlC:
				return
			}
		case termbox.EventError:
			return
		}
	}

	fmt.Printf("\nYou have selected: %s\n", options[selectedIndex])
}
