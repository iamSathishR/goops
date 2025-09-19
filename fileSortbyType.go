package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir := "." // (.) for current directory else replace for need

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			ext := filepath.Ext(path)
			if ext != "" {
				folder := filepath.Join(dir, ext[1:]) // removing "."
				os.MkdirAll(folder, os.ModePerm)
				newPath := filepath.Join(folder, info.Name())
				os.Rename(path, newPath)
				fmt.Printf("Moved; %s -> %s\n", path, newPath)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("error:", err)
	}
}
