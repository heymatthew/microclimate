package main_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	cmd "github.com/heymatthew/microclimate/cmd"
	"github.com/heymatthew/microclimate/pkg"
	"github.com/matryer/is"
	"golang.org/x/net/html"
)

func countTags(tag string, doc *html.Node) int {
	var count int
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		count += countTags(tag, c)
	}
	if doc.Type == html.ElementNode && doc.Data == tag {
		count++
	}
	return count
}

func TestCountTags(t *testing.T) {
	is := is.New(t)
	doc, err := html.Parse(strings.NewReader(`
		<html>
			<body>
				<a href="https://stuff.co.nz">stuff</a>
			</body>
		</html>
	`))
	is.NoErr(err)
	is.Equal(countTags("a", doc), 1)
}

func TestSetupCache(t *testing.T) {
	is := is.New(t)
	is.True(cmd.CacheDir() != "")
}

func TestSetupRouter(t *testing.T) {
	t.Run("root resolves", func(t *testing.T) {
		is := is.New(t)
		router := cmd.SetupRouter(pkg.Cache{})
		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/", nil)
		is.NoErr(err)
		router.ServeHTTP(w, req)
		is.Equal(200, w.Code)
	})

	t.Run("css path is sensible", func(t *testing.T) {
		is := is.New(t)
		router := cmd.SetupRouter(pkg.Cache{})
		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/static/style.css", nil)
		is.NoErr(err)
		router.ServeHTTP(w, req)
		is.Equal(200, w.Code)
	})

	t.Run("lists articles from cache", func(t *testing.T) {
		is := is.New(t)
		cache := pkg.Cache{
			Articles: []pkg.Sample{
				{Path: "foo"},
				{Path: "bar"},
			},
		}
		router := cmd.SetupRouter(cache)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)
		router.ServeHTTP(w, req)
		doc, err := html.Parse(w.Body)
		is.NoErr(err)
		is.Equal(countTags("a", doc), 2)
	})
}
