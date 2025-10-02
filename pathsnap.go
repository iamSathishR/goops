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

		if entry.IsDir() {
			line := fmt.Sprintf("%s %s %s/\n", prefix, connector, entry.Name())
			f.WriteString(line)
			newPrefix := prefix
			if i == len(entries)-1 {
				newPrefix += "    "
			} else {
				newPrefix += " |   "
			}
			printTree(filepath.Join(path, entry.Name()), newPrefix, f)
		} else {
			info, err := entry.Info()
			if err != nil {
				continue
			}
			line := fmt.Sprintf("%s %s %s (%s)\n", prefix, connector, entry.Name(), hrSize(info.Size()))
			f.WriteString(line)
		}
	}
	return nil
}

func hrSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(size)/float64(div), "KMGTPE"[exp])
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
		fmt.Println("Cannot create snapshot file:", err)
		return
	}
	defer f.Close()

	f.WriteString(dir + "\n")
	printTree(dir, "", f)

	fmt.Println("Directory snapshot saved to:", snapshotFile)
}
// C:\Users\PC\Downloads\GoOps
// |__ Folder1
// │  |__ one.go (2.3 KB)
// |__ Folder2
// │  |__ two.go (2.3 KB)
// |__ one.go (2.3 KB)
// |__ pathsnap.go (2.3 KB)
// |__ two.go (2.3 KB)
// |__ three.go (2.3 KB)
// |__ snapshot_20251001_233105.txt (20 KB)
