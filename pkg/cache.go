package pkg

import (
	"os"
	"path/filepath"
)

type Cache struct {
	Dir      string
	Articles []Sample
}

func (c *Cache) Load() error {
	return filepath.Walk(c.Dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		c.Articles = append(c.Articles, Sample{Path: path})
		return nil
	})
}
