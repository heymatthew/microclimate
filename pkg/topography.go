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

type Topography struct {
	Dir     string
	Samples []Sample
}

func (t *Topography) Load() error {
	return filepath.Walk(t.Dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		t.Samples = append(t.Samples, Sample{Path: path})
		return nil
	})
}
