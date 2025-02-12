package main

import (
	"fmt"
	"os"
	"bufio"
	"path/filepath"
	"os/exec"
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

func run(inputFilePath string) (string) {


	// Obtener la ruta completa sin la extensión
	filename := inputFilePath[:len(inputFilePath)-len(filepath.Ext(inputFilePath))]
	outputFile := filename + ".o"

	// Mostrar el nombre completo del archivo de salida
	content := "Object file " + outputFile + "\n"

	// Comando para ensamblar el archivo hello.s a hello.o
	cmd := exec.Command("as", "-o", outputFile, inputFilePath)

	// Ejecutar el comando y capturar cualquier error
	err := cmd.Run()
	if err != nil {
		// Si hay un error al ejecutar el comando, lo mostramos
		content += "Error executing as command.\n"
	} else {
		// Si todo es exitoso
		content += "Object file generated successfully!\n"
	}

	content += "Executable file " + filename + "\n"

	// Comando para ensamblar el archivo hello.s a hello.o
	cmd = exec.Command("ld", "-o", filename, outputFile)

	// Ejecutar el comando y capturar cualquier error
	err = cmd.Run()
	if err != nil {
		// Si hay un error al ejecutar el comando, lo mostramos
		content += "Error executing ld command.\n"
	} else {
		// Si todo es exitoso
		content += "Executable file generated successfully!\n"
	}

	// Comando para ensamblar el archivo hello.s a hello.o
	filename = "./" + filename
	// Ejecutar el comando y capturar cualquier error
	cmd = exec.Command(filename)

	output, err := cmd.CombinedOutput()  // CombinedOutput captura tanto stdout como stderr
	if err != nil {
		content += "Error reading Stdout.\n"
	}

	content += "\n=== Stdout ===\n"
	content += string(output)
	content += "\n==============\n"

	content += "Press any key to continue...\n"

	return content
}
