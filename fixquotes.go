package main

import (
	"regexp"
	"strings"
)

func fixQuotes(input string) string {
	// Регулярка: находим всё между одинарными кавычками
	re := regexp.MustCompile(`'([^']*?)'`)

	// Заменяем каждый матч, убирая пробелы по краям
	result := re.ReplaceAllStringFunc(input, func(match string) string {
		// match это что-то вроде "' awesome '"
		inner := match[1 : len(match)-1] // убираем кавычки
		inner = strings.TrimSpace(inner) // удаляем пробелы с краёв
		return "'" + inner + "'"         // возвращаем с кавычками
	})

	return result
}
