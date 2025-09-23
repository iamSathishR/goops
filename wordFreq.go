package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("to use; go run wordfreq.go samplefile")
		return
	}
	file := os.Args[1]

	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("error reading file: ", err)
		return
	}

	clean := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsSpace(r) {
			return unicode.ToLower(r)
		}
		return -1
	}, string(data))

	words := strings.Fields(clean)
	freq := make(map[string]int)

	for _, w := range words {
		freq[w]++
	}
	fmt.Println("Word frequencies:  ")
	for k, v := range freq {
		fmt.Printf("%s: %d\n", k, v)
	}
}
