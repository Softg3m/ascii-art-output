package main

import (
	"fmt"
)
func main(){
	args := os.Args[1:]

	if len(args) == 0 {
		Usage()
		return
	}

}

fumc Usage(){
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("EX: go run . --output=<fileName.txt> something standard")
}