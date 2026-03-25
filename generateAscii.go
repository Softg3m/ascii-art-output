package main

import (
	"fmt"
	"os"
	"strings"
)

func GenerateAscii(text string, banner string) string {
	data, err := os.ReadFile(banner + ".txt")
	if err != nil {
		fmt.Println("Error reading banner file", err)
	}

	lines := strings.Split(string(data), "\n")
	result := ""

	words := strings.Split(text, "\\n")

	for _, word := range words {
		for i := 1; i <= 8; i++ {
			for _, char := range word {
				asciiIndex := int(char) - 32
				start := asciiIndex * 9

				if start+i >= len(lines) {
					continue
				}

				result += lines[start+i]
			}
			result += "\n"
		}
	}

	return result
}
