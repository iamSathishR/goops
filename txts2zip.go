package main

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"time"
)

func main() {
	z, _ := os.Create("txt_backup_" + time.Now().Format("20060102_150405") + ".zip")
	defer z.Close()
	w := zip.NewWriter(z)
	defer w.Close()

	filepath.WalkDir(".", func(p string, d os.DirEntry, _ error) error {
		if d.IsDir() || filepath.Ext(d.Name()) != ".txt" {
			return nil
		}
		f, _ := os.Open(p)
		defer f.Close()
		r, _ := filepath.Rel(".", p)
		c, _ := w.Create(r)
		io.Copy(c, f)
		return nil
	})
}
