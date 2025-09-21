package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	var latestFile string
	var latestTime time.Time

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		if info.ModTime().After(latestTime) {
			latestTime = info.ModTime()
			latestFile = path
		}
		return nil
	})

	if latestFile == "" {
		fmt.Println("no files found")
	} else {
		fmt.Printf("latest modified file: %s\n modified at: %s\n", latestFile, latestTime.Format(time.RFC1123)) //std. HTTP date format
	}
}
