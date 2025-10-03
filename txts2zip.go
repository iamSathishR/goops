package main

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"time"
)

func main() {
	z, err := os.Create("txt_backup_" + time.Now().Format("20060102_150405") + ".zip")
	if err != nil {
		return
	}
	defer z.Close()
	w := zip.NewWriter(z)
	defer w.Close()

	filepath.WalkDir(".", func(p string, d os.DirEntry, e error) error {
		if e != nil || d.IsDir() || filepath.Ext(d.Name()) != ".txt" {
			return nil
		}
		f, err := os.Open(p)
		if err != nil {
			return nil
		}
		defer f.Close()
		rel, err := filepath.Rel(".", p)
		if err != nil {
			return nil
		}
		c, err := w.Create(rel)
		if err != nil {
			return nil
		}
		io.Copy(c, f)
		return nil
	})
}
