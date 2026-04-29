package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	text := string(data)

	// применяем все функции последовательно
	text = convert(text)
	text = fixPunct(text)
	text = fixQuotes(text)
	text = fixArticle(text)

	// записываем результат один раз
	err = os.WriteFile(outputFile, []byte(text), 0o644)
	if err != nil {
		panic(err)
	}
}
