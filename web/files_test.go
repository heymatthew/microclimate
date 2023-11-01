package web_test

import (
	"strings"
	"testing"

	web "github.com/heymatthew/microclimate/web"
	"github.com/matryer/is"
)

func TestFiles(t *testing.T) {
	t.Run("errors when file missing", func(t *testing.T) {
		is := is.New(t)
		_, err := web.Files.ReadFile("missing.txt")
		is.True(err != nil)
	})

	t.Run("finds static content", func(t *testing.T) {
		is := is.New(t)
		_, err := web.Files.ReadFile("static/style.css")
		is.True(err == nil)
	})

	t.Run("finds index", func(t *testing.T) {
		is := is.New(t)
		index, err := web.Files.ReadFile("templates/index.html.tmpl")
		is.True(err == nil)
		is.True(strings.HasPrefix(string(index), "<!doctype html>"))
	})
}
