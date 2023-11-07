package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	cmd "github.com/heymatthew/microclimate/cmd"
	"github.com/heymatthew/microclimate/pkg"
	"github.com/matryer/is"
)

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
}
