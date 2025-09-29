package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func main() {
	dir := "."
	var files []string

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	if len(files) > 0 {
		rand.Seed(time.Now().UnixNano())
    // seed for random unique results w. nanoseconds
		fmt.Println("random pick: ", files[rand.Intn(len(files))])
	} else {
		fmt.Println("no file found")
	}
}
