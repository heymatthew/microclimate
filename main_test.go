package microclimate_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/heymatthew/microclimate"
	"github.com/matryer/is"
)

func TestTopographyErrors(t *testing.T) {
	is := is.New(t)

	top := microclimate.Topography{Dir: "/made/up/folder"}
	err := top.Load()
	is.True(err != nil)
	is.Equal(len(top.Samples), 0)
}

func TestTopographyLoadFromDisk(t *testing.T) {
	is := is.New(t)

	// Create samples
	config := t.TempDir()
	for _, file := range []string{"foo.md", "bar.md"} {
		path := filepath.Join(config, file)
		err := os.WriteFile(path, []byte("hello world"), 0644)
		is.NoErr(err)
	}

	// Make sure they're both present
	top := microclimate.Topography{Dir: config}
	is.Equal(len(top.Samples), 0)
	top.Load()
	is.Equal(len(top.Samples), 2)
}
