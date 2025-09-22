package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("gowordydemo.txt")
	if err != nil {
		fmt.Println("error reading file: ", err)
		return
	}

	var flip []rune
	for _, ch := range string(data) {
		switch {
		case ch >= 'a' && ch <= 'z':
			flip = append(flip, ch-32)
		case ch >= 'A' && ch <= 'Z':
			flip = append(flip, ch+32)
		default:
			flip = append(flip, ch) // for leaving numbers/symbols as itis
		}
	}

	fmt.Println("original:")
	fmt.Println(string(data))
	fmt.Println()
	fmt.Println("<_><_><_><_><_><_><_><_><_><_> :P")
	fmt.Println("\n flipped:")
	fmt.Println(string(flip))
}
