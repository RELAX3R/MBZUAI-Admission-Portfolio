package main

import "regexp"

func fixArticle(input string) string {
	// Заменяем "a" на "an" перед гласной или 'h', при этом сохраняем исходные пробелы/переносы
	reLower := regexp.MustCompile(`\ba(\s+)([aeiouhAEIOUH])`)
	input = reLower.ReplaceAllString(input, "an$1$2")

	// То же для заглавной "A"
	reUpper := regexp.MustCompile(`\bA(\s+)([aeiouhAEIOUH])`)
	input = reUpper.ReplaceAllString(input, "An$1$2")

	return input
}
