package main

import (
	"github.com/nsf/termbox-go"
	"strings"
)

func isCtrl(word string) bool {
	return len(word) > 0 && word[0] == '^'
}

func renderText(content string, cursorX, cursorY int, clrScr bool) {
	x, y := 0, -1
	
	if clrScr {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	}
	
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		x = 0
		y++
		attr := termbox.ColorWhite
		for _, char := range line {
			if char == '.' {
				attr = termbox.ColorMagenta
			}
			if char == ' ' {
				attr = termbox.ColorWhite
			}
			termbox.SetCell(x, y, char, attr, termbox.ColorDefault)
			x++
		}
	}
	
	x = 0
	message := "^S Save   ^X Exit"
	words := strings.Fields(message)
	for _, word := range words {
		attr := termbox.ColorDefault
		if isCtrl(word) {
			attr = termbox.AttrReverse
			for _, char := range word {
				termbox.SetCell(x, 23, char, termbox.ColorWhite, attr)
				x++
			}
		} else {
			for _, char := range word {
				termbox.SetCell(x, 23, char, termbox.ColorWhite, attr)
				x++
			}
		}
		x++
	}
	termbox.SetCursor(cursorX, cursorY)
	termbox.Flush()
}
