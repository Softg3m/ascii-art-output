package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		Usage()
		return
	}
	
	// go run . --output=banner.txt "hello" standard
	if strings.HasPrefix(args[0], "--output=") {
		fileName := strings.TrimPrefix(args[0], "--output=")

		if fileName == "" || !strings.HasSuffix(fileName, ".txt") || len(args) < 2 {
			Usage()
			return
		}

		text := args[1]
		banner := "standard"

		if len(args) >= 3 {
			banner = args[2]
		}

		ascii := GenerateAscii(text, banner)

		err := os.WriteFile(fileName, []byte(ascii), 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			return
		}

	} else {
		//go run . "hello" standard
		text := args[0]
		banner := "standard"

		if len(args) >= 2 {
			banner = args[1]
		}

		ascii := GenerateAscii(text, banner)
		fmt.Print(ascii)
	}
}

func Usage() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("EX: go run . --output=<fileName.txt> something standard")
}

func GenerateAscii(text string, banner string) string {
	data, err := os.ReadFile(banner + ".txt")
	if err != nil {
		return "Error reading banner file\n"
	}

	lines := strings.Split(string(data), "\n")
	result := ""

	words := strings.Split(text, "\\n")

	for _, word := range words {
		for i := 0; i < 8; i++ {
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