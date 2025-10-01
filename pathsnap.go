package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func printTree(path string, prefix string, f *os.File) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for i, entry := range entries {
		connector := "|__ "
		if i == len(entries)-1 {
			connector = "|__ "
		}
		line := fmt.Sprintf("%s %s %s\n", prefix, connector, entry.Name())
		f.WriteString(line)

		if entry.IsDir() {
			newPrefix := prefix
			if i == len(entries)-1 {
				newPrefix += "  "
			} else {
				newPrefix += "│  "
			}
			printTree(filepath.Join(path, entry.Name()), newPrefix, f)
		}
	}
	return nil
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Cannot get current directory:", err)
		return
	}
	
	snapshotFile := fmt.Sprintf("snapshot_%s.txt", time.Now().Format("20060102_150405"))
	f, err := os.Create(snapshotFile)
	if err != nil {
		fmt.Println("cannot create snapshot file:", err)
		return
	}
	defer f.Close()
	
	f.WriteString(dir + "\n")
	printTree(dir, "", f)

	fmt.Println("Directory snapshot saved to: ", snapshotFile)
}

// C:\Users\PC\Downloads\GoOps
// |__ Folder1
// │  |__ one.go
// |__ Folder2
// │  |__ two.go
// |__ one.go
// |__ pathsnap.go
// |__ two.go
// |__ three.go
// |__ snapshot_20251001_233105.txt
