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

func TestTopography(t *testing.T) {
	t.Run("gracefully handles missing folders", func(t *testing.T) {
		is := is.New(t)

		top := pkg.Topography{Dir: "/made/up/folder"}
		err := top.Load()
		is.True(err != nil)
		is.Equal(len(top.Samples), 0)
	})

	t.Run("finds files on disk", func(t *testing.T) {
		is := is.New(t)

		// Create samples
		config := t.TempDir()
		for _, file := range []string{"foo.md", "bar.md"} {
			path := filepath.Join(config, file)
			err := os.WriteFile(path, []byte("hello world"), 0644)
			is.NoErr(err)
		}

		// Make sure they're both present
		top := pkg.Topography{Dir: config}
		is.Equal(len(top.Samples), 0)
		is.NoErr(top.Load())
		is.Equal(len(top.Samples), 2)
	})

	t.Run("loads content", func(t *testing.T) {
		is := is.New(t)

		// Create samples
		config := t.TempDir()
		for _, str := range []string{"aaa", "bbb"} {
			file := str + ".md"
			path := filepath.Join(config, file)
			err := os.WriteFile(path, []byte(str), 0644)
			is.NoErr(err)
		}

		top := pkg.Topography{Dir: config}
		is.NoErr(top.Load())
		is.True(strings.Contains(top.Samples[0].Content(), "aaa"))
		is.True(strings.Contains(top.Samples[1].Content(), "bbb"))
	})

	t.Run("translates markdown", func(t *testing.T) {
		is := is.New(t)

		// Create samples
		config := t.TempDir()
		str := dedent.Dedent(`
			# Heading

			content content content
		`)
		path := filepath.Join(config, "test.md")
		err := os.WriteFile(path, []byte(str), 0644)
		is.NoErr(err)

		top := pkg.Topography{Dir: config}
		is.NoErr(top.Load())
		body := top.Samples[0].Content()
		fmt.Println(body)
		is.True(strings.Contains(body, "content content content"))
		is.True(strings.Contains(body, "<h1>Heading</h1>"))
	})
}
