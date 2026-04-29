package main

import (
	"regexp"
	"strings"
)

func fixPunct(input string) string {
	// 1. Специальные группы
	input = strings.ReplaceAll(input, " ...", "...")
	input = strings.ReplaceAll(input, " !?", "!?")

	// 2. Удаляем пробел перед обычными знаками
	re := regexp.MustCompile(`\s+([.,!?;:])`)
	input = re.ReplaceAllString(input, "$1")

	// 3. Добавляем пробел после знака, если сразу идёт буква/цифра
	re2 := regexp.MustCompile(`([.,!?;:])([^\s.,!?;:])`)
	input = re2.ReplaceAllString(input, "$1 $2")

	// 4. Обрезаем лишние пробелы в начале/конце
	input = strings.TrimSpace(input)

	return input
}
