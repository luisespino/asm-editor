package main

import (
	"github.com/nsf/termbox-go"
	"strings"
)

func renderText(content string, cursorX, cursorY int, clrScr bool) {
	x, y := 0, -1
	
	if clrScr {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	}
	
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		x = 0
		y++
		for _, char := range line {
			termbox.SetCell(x, y, char, termbox.ColorWhite, termbox.ColorDefault)
			x++
		}
	}
	
	message := "Ctrl: Save   eXit"
	for i, char := range message {
		attr := termbox.ColorDefault
		if  char == 'S' || char == 'X' {
			attr = termbox.AttrReverse
		}
		termbox.SetCell(i, 23, char, termbox.ColorWhite, attr)
	}
	termbox.SetCursor(cursorX, cursorY)
	termbox.Flush()
}
