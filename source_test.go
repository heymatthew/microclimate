package main

import (
	"testing"
)

type test struct { input string }

var validUrls = []test {
	{ "https://test.net" },
	{ "http://test.net" },
}

func TestValid(t *testing.T) {
	for _, tt := range validUrls {
		s := Source{
			headline: "Otters",
			url: tt.input,
		}
		if !s.Valid() {
			t.Errorf("Expected '%s' to be a valid url", tt.input)
		}
	}
}

var invalidUrls = []test {
	{ "" },
	{ "hhttp://test.net" },
}

func TestInvalid(t *testing.T) {
	for _, tt := range invalidUrls {
		s := Source{
			headline: "Otters",
			url: tt.input,
		}
		if s.Valid() {
			t.Errorf("Expected '%s' to be invalid", tt.input)
		}
	}
}
