package main

import (
	"github.com/nsf/termbox-go"
	"strings"
)

func isCtrl(word string) bool {
	return len(word) > 0 && word[0] == '^'
}

func setColor(word string) termbox.Attribute {
	if len(word) > 0 && word[0] == '.' {
		return termbox.ColorLightMagenta
	}
	if len(word) > 0 && word[len(word)-1] == ':' {
		return termbox.ColorLightBlue
	}
	if len(word) > 0 && word[0] == '"' {
		return termbox.ColorLightYellow
	}
	if len(word) > 0 && word[len(word)-1] == '"' {
		return termbox.ColorLightYellow
	}
	return termbox.ColorWhite
}


func renderText(content string, cursorX, cursorY int, clrScr bool) {
	x, y := 0, 0
	
	if clrScr {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	}
	
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		x = 0
		y++

		word := ""

		attr := termbox.ColorWhite
		for _, char := range line {
			if char == ' ' || char == '\t' {
				if word != "" {
					attr = setColor(word)
					for _, wChar := range word {
						termbox.SetCell(x, y, wChar, attr, termbox.ColorDefault)
						x++
					}
					word = ""	
				}
				termbox.SetCell(x, y, char, attr, termbox.ColorDefault)
				x++
			} else {
				word += string(char)
			}
			
		}

		if word != "" {
			attr = setColor(word)
			for _, wChar := range word {
				termbox.SetCell(x, y, wChar, attr, termbox.ColorDefault)
				x++
			}
			word = ""	
		}

	}

	programName := "   asm-editor"
	centerPos := (80 - len(filename)) / 2
	for x := 0; x < 80; x++ {
		termbox.SetCell(x, 0, ' ', termbox.ColorBlack, termbox.ColorWhite)
	}
	for i, char := range programName {
		termbox.SetCell(i, 0, char, termbox.ColorBlack, termbox.ColorWhite)
	}
	for i, char := range filename {
		termbox.SetCell(centerPos+i, 0, char, termbox.ColorBlack, termbox.ColorWhite)
	}

	x = 0
	message := "^S Save   ^X Exit   ^R Run"
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
