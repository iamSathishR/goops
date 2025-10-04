package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root, err := os.Getwd()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			entries, _ := os.ReadDir(path)
			if len(entries) == 0 {
				fmt.Println("empty directory:", path)
			}
		}
		return nil
	})
}
