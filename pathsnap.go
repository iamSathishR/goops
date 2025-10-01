package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir, _ := os.Getwd()
	snapshotFile := filepath.Join(dir, "snapshot.txt")

	file, err := os.Create(snapshotFile)
	if err != nil {
		fmt.Println("error creating snapshot file:", err)
		return
	}
	defer file.Close()

	filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if !d.IsDir() {
			info, e := d.Info()
			if e == nil {
				fmt.Fprintf(file, "%s | Size: %d bytes | Modified: %s\n",
					path, info.Size(), info.ModTime().Format("2006-01-02 15:04:05"))
			}
		}
		return nil
	})

	fmt.Println("Snapshot saved to:", snapshotFile)
}

// sampleline: C:\Users\PC\Downloads\folder\SQL-cheat-sheet.pdf | Size: 230259 bytes | Modified: 2025-09-30 18:04:24
