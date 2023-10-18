package pkg

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/yuin/goldmark"
)

type Sample struct {
	Path string
}

func (s Sample) Content() string {
	content, err := os.ReadFile(s.Path)
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	err = goldmark.Convert(content, &buf)
	if err != nil {
		panic(err)
	}
	return string(buf.String())
}

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