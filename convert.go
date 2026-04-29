package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func convert(input string) string {
	// ловим либо обычные "слова" без скобок, либо целиком скобочные функции с любыми пробелами внутри
	re := regexp.MustCompile(`([^\s()]+)|\(\s*(?:hex|bin|up(?:\s*,\s*\d+)?|low(?:\s*,\s*\d+)?|cap(?:\s*,\s*\d+)?)\s*\)`)

	// обрабатываем построчно, чтобы сохранить переводы строк
	lines := strings.Split(input, "\n")
	out := make([]string, 0, len(lines))

	for _, line := range lines {
		words := re.FindAllString(line, -1)

		// нормализуем скобочные токены: убираем пробелы внутри, чтобы получить, например, "(up,3)"
		for i, w := range words {
			if strings.HasPrefix(w, "(") && strings.HasSuffix(w, ")") {
				words[i] = strings.ReplaceAll(w, " ", "")
			}
		}

		for i := 1; i < len(words); i++ {
			switch {
			case words[i] == "(hex)":
				if i-1 >= 0 {
					raw := words[i-1]
					s := strings.Trim(raw, `"'`) // убираем внешние кавычки для парсинга
					value, err := strconv.ParseInt(s, 16, 64)
					if err == nil {
						newVal := fmt.Sprintf("%d", value)
						words[i-1] = strings.Replace(raw, s, newVal, 1) // сохраняем кавычки
						words = append(words[:i], words[i+1:]...)
						i--
					}
				}
			case words[i] == "(bin)":
				if i-1 >= 0 {
					raw := words[i-1]
					s := strings.Trim(raw, `"'`)
					value, err := strconv.ParseInt(s, 2, 64)
					if err == nil {
						newVal := fmt.Sprintf("%d", value)
						words[i-1] = strings.Replace(raw, s, newVal, 1)
						words = append(words[:i], words[i+1:]...)
						i--
					}
				}
			case words[i] == "(up)":
				if i-1 >= 0 {
					words[i-1] = strings.ToUpper(words[i-1])
					words = append(words[:i], words[i+1:]...)
					i--
				}
			case strings.HasPrefix(words[i], "(up,"):
				if i-1 >= 0 {
					var n int
					if _, err := fmt.Sscanf(words[i], "(up,%d)", &n); err == nil {
						for j := 0; j < n; j++ {
							idx := i - 1 - j
							if idx >= 0 {
								words[idx] = strings.ToUpper(words[idx])
							}
						}
						words = append(words[:i], words[i+1:]...)
						i--
					}
				}
			case words[i] == "(low)":
				if i-1 >= 0 {
					words[i-1] = strings.ToLower(words[i-1])
					words = append(words[:i], words[i+1:]...)
					i--
				}
			case strings.HasPrefix(words[i], "(low,"):
				if i-1 >= 0 {
					var n int
					if _, err := fmt.Sscanf(words[i], "(low,%d)", &n); err == nil {
						for j := 0; j < n; j++ {
							idx := i - 1 - j
							if idx >= 0 {
								words[idx] = strings.ToLower(words[idx])
							}
						}
						words = append(words[:i], words[i+1:]...)
						i--
					}
				}
			case words[i] == "(cap)":
				if i-1 >= 0 {
					word := words[i-1]
					runes := []rune(word)
					if len(runes) > 0 {
						runes[0] = unicode.ToUpper(runes[0])
						words[i-1] = string(runes)
					}
					words = append(words[:i], words[i+1:]...)
					i--
				}
			case strings.HasPrefix(words[i], "(cap,"):
				if i-1 >= 0 {
					var n int
					if _, err := fmt.Sscanf(words[i], "(cap,%d)", &n); err == nil {
						for j := 0; j < n; j++ {
							idx := i - 1 - j
							if idx >= 0 {
								runes := []rune(words[idx])
								if len(runes) > 0 {
									runes[0] = unicode.ToUpper(runes[0])
									words[idx] = string(runes)
								}
							}
						}
						words = append(words[:i], words[i+1:]...)
						i--
					}
				}
			}
		}
		out = append(out, strings.Join(words, " "))
	}

	return strings.Join(out, "\n")
}
