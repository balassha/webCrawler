package main

import (
	"htmlparser/infrastructure/httpClient"
	"htmlparser/webscrapper"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestGetHTMLVersion(t *testing.T) {
	cases := []struct {
		html     string
		expected string
	}{
		{
			html:     `"-//W3C//DTD HTML 4.01//EN"`,
			expected: "HTML 4.01 Strict",
		},
		{
			html:     `"-//W3C//DTD HTML 4.01 Transitional//EN"`,
			expected: "HTML 4.01 Transitional",
		},
		{
			html:     `"-//W3C//DTD HTML 4.01 Frameset//EN"`,
			expected: "HTML 4.01 Frameset",
		},
		{
			html:     `"-//W3C//DTD XHTML 1.0 Strict//EN"`,
			expected: "XHTML 1.0 Strict",
		},
		{
			html:     `"-//W3C//DTD XHTML 1.0 Transitional//EN"`,
			expected: "XHTML 1.0 Transitional",
		},
		{
			html:     `"-//W3C//DTD XHTML 1.0 Frameset//EN"`,
			expected: "XHTML 1.0 Frameset",
		},
		{
			html:     `"-//W3C//DTD XHTML 1.1//EN"`,
			expected: "XHTML 1.1",
		},
		{
			html:     `"<!DOCTYPE html>"`,
			expected: "HTML 5",
		},
		{
			html:     "",
			expected: "UNKNOWN",
		},
		{
			html:     "1",
			expected: "UNKNOWN",
		},
		{
			html:     "test",
			expected: "UNKNOWN",
		},
	}
	for _, c := range cases {
		have := webscrapper.GetHTMLVersion(c.html)
		assert.Equal(t, c.expected, have)
	}
}

func TestProcessInaccessibleLinks(t *testing.T) {
	cases := []struct {
		links    []string
		expected int
	}{
		{
			links:    []string{"http://"},
			expected: 1,
		},
		{
			links:    []string{"https://"},
			expected: 1,
		},
		{
			links:    []string{"https://com"},
			expected: 1,
		},
		{
			links:    []string{"https://www.google.com"},
			expected: 0,
		},
		{
			links:    []string{"https://www.google.com", "https://www.google.com"},
			expected: 0,
		},
		{
			links:    []string{"https://www.google.co.in", "https://www.google.com"},
			expected: 0,
		},
	}
	var client httpClient.HttpClient
	client.Initialize()
	for _, c := range cases {
		have := webscrapper.ProcessInaccessibleLinks(&client, c.links)
		assert.Equal(t, c.expected, len(have))
	}
}

func TestGetInternalAndExternalLinks(t *testing.T) {
	var scrapper webscrapper.Scrapper
	cases := []struct {
		links    []string
		expected []int
	}{
		{
			links:    []string{"", ""},
			expected: []int{0, 2},
		},
		{
			links:    []string{"https://www.google.co.in", "/preferences?hl=en"},
			expected: []int{1, 1},
		},
	}
	for _, c := range cases {
		scrapper.Links = c.links
		have, _ := scrapper.GetInternalAndExternalLinks()
		assert.Equal(t, c.expected[0], len(have.ExternalLinks))
		assert.Equal(t, c.expected[1], len(have.InternalLinks))

	}
}

func TestSliceContains(t *testing.T) {
	cases := []struct {
		slice    []string
		value    string
		expected bool
	}{
		{
			slice:    []string{""},
			value:    "",
			expected: true,
		},
		{
			slice:    []string{"1"},
			value:    "1",
			expected: true,
		},
		{
			slice:    []string{"1"},
			value:    "test",
			expected: false,
		},
	}

	for _, c := range cases {
		have := webscrapper.SliceContains(c.slice, c.value)
		assert.Equal(t, c.expected, have)
	}
}
