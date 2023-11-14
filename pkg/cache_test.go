package pkg_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/heymatthew/microclimate/pkg"
	"github.com/matryer/is"
)

func createSamples(dir string, names []string) error {
	for _, name := range names {
		path := filepath.Join(dir, name)
		err := os.WriteFile(path, []byte(name), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func TestCache(t *testing.T) {
	t.Run("gracefully handles missing folders", func(t *testing.T) {
		is := is.New(t)
		cache := pkg.Cache{Dir: "/made/up/folder"}
		err := cache.Load()
		is.True(err != nil)
		is.Equal(len(cache.Articles), 0)
	})

	t.Run("loads content from disk", func(t *testing.T) {
		is := is.New(t)
		dir := t.TempDir()
		is.NoErr(createSamples(dir, []string{"aaa.md", "bbb.md"}))
		cache := pkg.Cache{Dir: dir}
		is.Equal(len(cache.Articles), 0)
		is.NoErr(cache.Load())
		is.Equal(len(cache.Articles), 2)
		is.True(strings.Contains(cache.Articles[0].Content(), "aaa"))
		is.True(strings.Contains(cache.Articles[1].Content(), "bbb"))
	})
}
