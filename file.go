package main

import (
	"fmt"
	"os"
	"bufio"
	"path/filepath"
)


func saveToFile(filename string, content string) {
	

	dirPath := filepath.Dir(filename)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		fmt.Println("Error making directories:", err)
		return
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	writer.Flush()
}

func loadFromFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var content string
	scanner := bufio.NewScanner(file)
	firstLine := true
	for scanner.Scan() {
    		if !firstLine {
        		content += "\n"
    		}
    		content += scanner.Text()
    		firstLine = false
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return content, nil
}