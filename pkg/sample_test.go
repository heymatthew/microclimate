package pkg_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/heymatthew/microclimate/pkg"
	"github.com/lithammer/dedent"
	"github.com/matryer/is"
)

func TestSample(t *testing.T) {
	t.Run("translates markdown", func(t *testing.T) {
		is := is.New(t)

		// Create samples
		str := dedent.Dedent(`
			# Heading

			content content content
		`)
		path := filepath.Join(t.TempDir() + "test.md")
		err := os.WriteFile(path, []byte(str), 0644)
		is.NoErr(err)

		sample := pkg.Sample{Path: path}
		body := sample.Content()
		fmt.Println(body)
		is.True(strings.Contains(body, "content content content"))
		is.True(strings.Contains(body, "<h1>Heading</h1>"))
	})
}
