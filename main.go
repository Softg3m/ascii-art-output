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
