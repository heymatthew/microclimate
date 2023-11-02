package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	cmd "github.com/heymatthew/microclimate/cmd"
	"github.com/matryer/is"
)

func TestSetupTopography(t *testing.T) {
	is := is.New(t)
	is.True(cmd.CacheDir() != "")
}

func TestSetupRouter(t *testing.T) {
	t.Run("root resolves", func(t *testing.T) {
		is := is.New(t)
		router := cmd.SetupRouter()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)
		router.ServeHTTP(w, req)
		is.Equal(200, w.Code)
	})

	t.Run("css path is sensible", func(t *testing.T) {
		is := is.New(t)
		router := cmd.SetupRouter()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/static/style.css", nil)
		router.ServeHTTP(w, req)
		is.Equal(200, w.Code)
	})
}
