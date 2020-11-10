package model

import (
	"net/url"
	"regexp"
)

// Shortcut allocates a path to a URL which the path will be redirected to
type Shortcut struct {
	Path string
	URL  string
}

// IsValid checks the consistency of a shortcut struct
func (shortcut *Shortcut) IsValid() bool {
	if shortcut.Path == "" || shortcut.URL == "" {
		return false
	}
	if len(shortcut.Path) > 50 {
		return false
	}
	matched, _ := regexp.MatchString("^[A-Za-z0-9-_]*$", shortcut.Path)
	if !matched {
		return false
	}
	_, err := url.ParseRequestURI(shortcut.URL)
	if err != nil {
		return false
	}
	return true
}
