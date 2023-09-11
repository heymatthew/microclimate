package main_test

import (
	"testing"

	cmd "github.com/heymatthew/microclimate/cmd"
	"github.com/matryer/is"
)

func TestSetupTopography(t *testing.T) {
	is := is.New(t)
	is.True(cmd.CacheDir() != "")
}
