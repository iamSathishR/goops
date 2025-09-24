package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("to use; go run groupByExt.go dir")
		return
	}

	root := os.Args[1]
	extCount := make(map[string]int)

	filepath.WalkDir(root, func(_ string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if !d.IsDir() {
			ext := strings.ToLower(filepath.Ext(d.Name()))
			if ext == "" {
				ext = "[no extension]"
			}
			extCount[ext]++
		}
		return nil
	})

	type kv struct {
		Ext   string
		Count int
	}
	var sorted []kv
	for k, v := range extCount {
		sorted = append(sorted, kv{k, v})
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Count > sorted[j].Count
	})

	fmt.Println("file type summary in:", root)
	for _, kv := range sorted {
		fmt.Printf("%s : %d files\n", kv.Ext, kv.Count)
	}
}
