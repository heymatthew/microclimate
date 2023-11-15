package pkg

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
)

type Cache struct {
	Dir      string
	Articles []Article
}

type Article struct {
	Path string
}

func (s Article) Content() string {
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

func (c *Cache) Load() error {
	return filepath.Walk(c.Dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".md") {
			return nil
		}
		c.Articles = append(c.Articles, Article{Path: path})
		return nil
	})
}
