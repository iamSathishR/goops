package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("to use: go run main.go <directory>")
		// run w/ ~~ go run filename.go "C:\directory\dude\here"
		return
	}
	dir := os.Args[1]
	files := make(map[int64][]string)

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() {
			size := info.Size()
			files[size] = append(files[size], path)
		}
		return nil
	})
	for size, paths := range files {
		if len(paths) > 1 {
			fmt.Println("possible duplicates (size:", size, "bytes):")
			for _, p := range paths {
				fmt.Println("  -", p)
			}
		}
	}
}
