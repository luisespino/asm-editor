package main

import "strings"

func backspace(content string, cursorX, cursorY int) (string, int, int) {
	lines := strings.Split(content, "\n")

	if cursorY > 0 || cursorX > 0 {
		if cursorX > 0 {
			line := lines[cursorY]
			line = line[:cursorX-1] + line[cursorX:]
			lines[cursorY] = line
			cursorX--
		} else if cursorY > 0 {
			cursorY--
			cursorX = len(lines[cursorY])
			lines = append(lines[:cursorY+1], lines[cursorY+2:]...)
		}
	}
	content = strings.Join(lines, "\n")
	return content, cursorX, cursorY
}

func enter(content string, cursorX, cursorY int) (string, int, int) {
    lines := strings.Split(content, "\n")
    
    currentLine := lines[cursorY]
    lines[cursorY] = currentLine[:cursorX]

    if len(currentLine[cursorX:]) > 0 {
        lines = append(lines[:cursorY+1], lines[cursorY:]...)
        lines[cursorY+1] = currentLine[cursorX:]
    } else {
        lines = append(lines[:cursorY+1], lines[cursorY:]...)
        lines[cursorY+1] = "" 
	}

    content = strings.Join(lines, "\n")
    cursorX = 0
    cursorY++ 

    return content, cursorX, cursorY
}



func delete(content string, cursorX, cursorY int) (string, int, int) {

	lines := strings.Split(content, "\n")


	if cursorY < len(lines) {
		line := lines[cursorY]


		if cursorX < len(line) {
			line = line[:cursorX] + line[cursorX+1:]
			lines[cursorY] = line
		} else if cursorY < len(lines)-1 {
			lines[cursorY] = lines[cursorY] + lines[cursorY+1] // Unir la línea actual con la siguiente
			lines = append(lines[:cursorY+1], lines[cursorY+2:]...) // Eliminar la siguiente línea
		}
	}

	content = strings.Join(lines, "\n")

	return content, cursorX, cursorY
}




func moveUp(content string, cursorX, cursorY int) (int, int) {
	lines := strings.Split(content, "\n")
	if cursorY > 0 {
		cursorY--
		if cursorX > len(lines[cursorY]) {
			cursorX = len(lines[cursorY]) 
		}
	}
	return cursorX, cursorY
}

func moveDown(content string, cursorX, cursorY int) (int, int) {
	lines := strings.Split(content, "\n")
	if cursorY < len(lines)-1 {
		cursorY++
		if cursorX > len(lines[cursorY]) {
			cursorX = len(lines[cursorY])
		}
	}
	return cursorX, cursorY
}

func moveLeft(content string, cursorX, cursorY int) (int, int) {
	if cursorX > 0 {
		cursorX--
	}
	return cursorX, cursorY
}

func moveRight(content string, cursorX, cursorY int) (int, int) {
	lines := strings.Split(content, "\n")
	if cursorX < len(lines[cursorY]) {
		cursorX++
	}
	return cursorX, cursorY
}
