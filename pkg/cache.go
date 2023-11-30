package pkg

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/text"
	"gopkg.in/yaml.v2"
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
		panic(fmt.Errorf("Can not read %s: %w", s.Path, err))
	}
	var buf bytes.Buffer
	err = goldmark.Convert(content, &buf)
	if err != nil {
		panic(err)
	}
	return string(buf.String())
}

// n.b. Does not handle tabs
var data = `
title: Easy!
`

type Metadata struct {
	Title string
}

func NewMetadata() Metadata {
	m := Metadata{}
	yaml.Unmarshal([]byte(data), &m)
	return m
}

func (s Article) Title() string {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			meta.New(
				meta.WithStoresInDocument(),
			),
		),
	)
	source := `---
title: toddler tantrums inline
---
data data data
`
	source = "Hello"

	document := markdown.Parser().Parse(text.NewReader([]byte(source)))
	metaData := document.OwnerDocument().Meta()
	title := metaData["title"]
	return title.(string)
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
