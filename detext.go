package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && !strings.Contains(info.Name(), ".") {
			fmt.Println("files w/ no extension:", path)
		}
		return nil
	})
}
