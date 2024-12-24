package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	fileMap := make(map[string]map[string]bool) // Mapa para armazenar em quais arquivos a linha aparece

	if len(files) == 0 {
		countLines(os.Stdin, counts, fileMap, "<stdin>")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileMap, arg)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 { // Se a linha aparece mais de uma vez
			fmt.Printf("Aparece(%d) vezes\tO caractere(%s)\n", n, line)
			fmt.Print("Nos arquivos: ")
			// Exibe os arquivos em que a linha apareceu
			for file := range fileMap[line] {
				fmt.Printf("%s ", file)
			}
			fmt.Println()
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileMap map[string]map[string]bool, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		// Adiciona o arquivo no mapa de arquivos que contém a linha
		if fileMap[line] == nil {
			fileMap[line] = make(map[string]bool)
		}
		fileMap[line][filename] = true
	}
	// NOTE: ignoring potential errors from input.Err()
}
