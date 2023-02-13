package main

import "regexp"

type Source struct {
	url, headline string
	priority      uint
}

var re = regexp.MustCompile("^https?://")

func (s Source) Valid() bool {
	if !re.MatchString(s.url) {
		return false
	}
	return true
}
