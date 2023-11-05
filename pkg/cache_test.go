package pkg_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/heymatthew/microclimate/pkg"
	"github.com/matryer/is"
)

func TestCache(t *testing.T) {
	t.Run("gracefully handles missing folders", func(t *testing.T) {
		is := is.New(t)
		cache := pkg.Cache{Dir: "/made/up/folder"}
		err := cache.Load()
		is.True(err != nil)
		is.Equal(len(cache.Articles), 0)
	})

	t.Run("finds files on disk", func(t *testing.T) {
		is := is.New(t)

		// Create samples
		dir := t.TempDir()
		for _, file := range []string{"foo.md", "bar.md"} {
			path := filepath.Join(dir, file)
			err := os.WriteFile(path, []byte("hello world"), 0644)
			is.NoErr(err)
		}

		// Make sure they're both present
		cache := pkg.Cache{Dir: dir}
		is.Equal(len(cache.Articles), 0)
		is.NoErr(cache.Load())
		is.Equal(len(cache.Articles), 2)
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

		cache := pkg.Cache{Dir: config}
		is.NoErr(cache.Load())
		is.True(strings.Contains(cache.Articles[0].Content(), "aaa"))
		is.True(strings.Contains(cache.Articles[1].Content(), "bbb"))
	})
}
