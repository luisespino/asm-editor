package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
	"strings"
)

var filename string

func main() {
	// capture filePath in args
	args := os.Args
	filename = args[1]

	// termbox initialization
	err := termbox.Init()
	if err != nil {
		fmt.Println("Error initializing termbox:", err)
		return
	}
	defer termbox.Close()

	// load file
	content, err := loadFromFile(filename)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("Error loading file:", err)
		return
	}

	// intial render
	cursorX, cursorY := 0, 1
	clrScr := true
	renderText(content, cursorX, cursorY, clrScr)

	// get user input
	for {
		switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				clrScr = false
				switch ev.Key {
					case termbox.KeyCtrlS:
						saveToFile(filename, content)
					case termbox.KeyCtrlX:
						return
					case termbox.KeyCtrlR:
						temp := run(filename)
						clrScr = true
						renderText(temp, cursorX, cursorY, clrScr)
						termbox.PollEvent()
					case termbox.KeyEnter:
						content, cursorX, cursorY = enter(content, cursorX, cursorY)
						clrScr = true
					case termbox.KeyTab:
						lines := strings.Split(content, "\n")
						lines[cursorY-1] = lines[cursorY-1][:cursorX] + "    " + lines[cursorY-1][cursorX:]
						content = strings.Join(lines, "\n")
						cursorX += 4
					case termbox.KeyBackspace, termbox.KeyBackspace2:
						content, cursorX, cursorY = backspace(content, cursorX, cursorY)
						clrScr = true
					case termbox.KeyDelete:
						content, cursorX, cursorY = delete(content, cursorX, cursorY)
						clrScr = true
					case termbox.KeyArrowUp:
						cursorX, cursorY = moveUp(content, cursorX, cursorY)
					case termbox.KeyArrowDown:
						cursorX, cursorY = moveDown(content, cursorX, cursorY)
					case termbox.KeyArrowLeft:
						cursorX, cursorY = moveLeft(content, cursorX, cursorY)
					case termbox.KeyArrowRight:
						cursorX, cursorY = moveRight(content, cursorX, cursorY)
					default:
						if ev.Key == termbox.KeySpace {
							lines := strings.Split(content, "\n")
							lines[cursorY-1] = lines[cursorY-1][:cursorX] + " " + lines[cursorY-1][cursorX:]
							content = strings.Join(lines, "\n")
						} else if ev.Ch != 0 {
							lines := strings.Split(content, "\n")
							lines[cursorY-1] = lines[cursorY-1][:cursorX] + string(ev.Ch) + lines[cursorY-1][cursorX:]
							content = strings.Join(lines, "\n")
						}
						cursorX++
				}
				renderText(content, cursorX, cursorY, clrScr)
		}
	}
}

