package pkg

import (
	"bytes"
	"os"

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
