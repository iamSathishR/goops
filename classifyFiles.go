package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	sizes := map[string][]string{"small(<1MB)": {}, "medium(1-10MB)": {}, "large(>10MB)": {}}

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size := info.Size()
			switch {
			case size < 1<<20:
				sizes["small(<1MB)"] = append(sizes["small(<1MB)"], path)
			case size < 10<<20:
				sizes["medium(1-10MB)"] = append(sizes["medium(1-10MB)"], path)
			default:
				sizes["large(>10MB)"] = append(sizes["large(>10MB)"], path)
			}
		}
		return nil
	})

	for k, v := range sizes {
		fmt.Println(k, len(v), "files")
	}
}
