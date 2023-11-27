package pkg_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/heymatthew/microclimate/pkg"
	"github.com/lithammer/dedent"
	"github.com/matryer/is"
)

func createArticles(dir string, names []string) error {
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
		is.NoErr(createArticles(dir, []string{"aaa.md", "bbb.md"}))
		cache := pkg.Cache{Dir: dir}
		is.Equal(len(cache.Articles), 0)
		is.NoErr(cache.Load())
		is.Equal(len(cache.Articles), 2)
		is.True(strings.Contains(cache.Articles[0].Content(), "aaa"))
		is.True(strings.Contains(cache.Articles[1].Content(), "bbb"))
	})

	t.Run("lists markdown files", func(t *testing.T) {
		is := is.New(t)
		dir := t.TempDir()
		is.NoErr(createArticles(dir, []string{"markdown.md"}))
		cache := pkg.Cache{Dir: dir}
		is.NoErr(cache.Load())
		is.Equal(len(cache.Articles), 1)
	})

	t.Run("does not list excluded files", func(t *testing.T) {
		excludes_list := []string{".gitignore", "guff.html"}
		is := is.New(t)
		dir := t.TempDir()
		is.NoErr(createArticles(dir, excludes_list))
		cache := pkg.Cache{Dir: dir}
		is.NoErr(cache.Load())
		is.Equal(len(cache.Articles), 0)
	})
}

func TestArticle(t *testing.T) {
	t.Run("translates markdown", func(t *testing.T) {
		is := is.New(t)
		sample := dedent.Dedent(`
			# Heading

			content content content
		`)
		path := filepath.Join(t.TempDir() + "test.md")
		err := os.WriteFile(path, []byte(sample), 0644)
		is.NoErr(err)

		article := pkg.Article{Path: path}
		body := article.Content()
		is.True(strings.Contains(body, "content content content"))
		is.True(strings.Contains(body, "<h1>Heading</h1>"))
	})
}
