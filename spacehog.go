package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func dirSize(path string) int64 {
	var total int64
	filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() {
			total += info.Size()
		}
		return nil
	})
	return total
}

func main() {
	root := `D:\` // select drive
  //root := "D:\\"
	threshold := int64(500 * 1024 * 1024) // 500 MB

	filepath.Walk(root, func(p string, info os.FileInfo, err error) error {
		if err != nil || !info.IsDir() {
			return nil
		}
		size := dirSize(p)
		if size > threshold {
			fmt.Printf("  Folder %s uses %.2f MB\n", p, float64(size)/1024/1024)
		}
		return nil
	})
}
