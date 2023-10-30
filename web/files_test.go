package web_test

import (
	"fmt"
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
		fmt.Println("Err is", err)
	})

	t.Run("can find standard templates", func(t *testing.T) {
		is := is.New(t)

		index, err := web.Files.ReadFile("templates/index.html.tmpl")
		fmt.Println("Err is", err)
		is.True(err == nil)
		is.True(strings.HasPrefix(string(index), "<!doctype html>"))
	})
}
