package main

import "strings"

func backspace(content string, cursorX, cursorY int) (string, int, int) {
    lines := strings.Split(content, "\n")

    if cursorY > 1 || cursorX > 0 {
        if cursorX > 0 {
            // El cursor no está al principio de la línea
            line := lines[cursorY-1]
            line = line[:cursorX-1] + line[cursorX:]
            lines[cursorY-1] = line
            cursorX--
        } else if cursorY > 1 {
            // El cursor está al principio de la línea
            cursorY-- // Mover el cursor a la línea anterior
            cursorX = len(lines[cursorY-1]) // Colocar el cursor al final de la línea anterior
            // Concatenar la línea actual con la línea anterior
            lines[cursorY-1] = lines[cursorY-1] + lines[cursorY]
            // Eliminar la línea vacía (ahora combinada)
            lines = append(lines[:cursorY], lines[cursorY+1:]...)
        }
    }

    content = strings.Join(lines, "\n")
    return content, cursorX, cursorY
}


func enter(content string, cursorX, cursorY int) (string, int, int) {
    lines := strings.Split(content, "\n")
    
    currentLine := lines[cursorY-1]
    lines[cursorY-1] = currentLine[:cursorX]

    if len(currentLine[cursorX:]) > 0 {
        lines = append(lines[:cursorY], lines[cursorY-1:]...)
        lines[cursorY] = currentLine[cursorX:]
    } else {
        lines = append(lines[:cursorY], lines[cursorY-1:]...)
        lines[cursorY] = "" 
	}

    content = strings.Join(lines, "\n")
    cursorX = 0
    cursorY++ 

    return content, cursorX, cursorY
}



func delete(content string, cursorX, cursorY int) (string, int, int) {

	lines := strings.Split(content, "\n")


	if cursorY-1 < len(lines) {
		line := lines[cursorY-1]


		if cursorX < len(line) {
			line = line[:cursorX] + line[cursorX+1:]
			lines[cursorY-1] = line
		} else if cursorY < len(lines) {
			lines[cursorY-1] = lines[cursorY-1] + lines[cursorY] // Unir la línea actual con la siguiente
			lines = append(lines[:cursorY], lines[cursorY+1:]...) // Eliminar la siguiente línea
		}
	}

	content = strings.Join(lines, "\n")

	return content, cursorX, cursorY
}




func moveUp(content string, cursorX, cursorY int) (int, int) {
	lines := strings.Split(content, "\n")
	if cursorY > 1 {
		cursorY--
		if cursorX > len(lines[cursorY-1]) {
			cursorX = len(lines[cursorY-1]) 
		}
	}
	return cursorX, cursorY
}

func moveDown(content string, cursorX, cursorY int) (int, int) {
	lines := strings.Split(content, "\n")
	if cursorY < len(lines) {
		cursorY++
		if cursorX > len(lines[cursorY-1]) {
			cursorX = len(lines[cursorY-1])
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
	if cursorX < len(lines[cursorY-1]) {
		cursorX++
	}
	return cursorX, cursorY
}
