package webscrapper

import (
	"fmt"
	"htmlparser/infrastructure/httpClient"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

//Get HTML version from html
func GetHTMLVersion(html string) string {
	var version = "UNKNOWN"
	doctypes := InitializeDoctypes()
	for doctype, matcher := range doctypes {
		match := strings.Contains(html, strings.ToLower(matcher)) || strings.Contains(html, matcher)
		if match {
			version = doctype
		}
	}

	return version
}

//Get Page title from html
func GetPageTitle(node *html.Node) string {
	var title string
	if node.Type == html.ElementNode && node.Data == "title" {
		return node.FirstChild.Data
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		title = GetPageTitle(c)
		if title != "" {
			break
		}
	}
	return title
}

// Range through all external links and identify unaccessible links
func ProcessInaccessibleLinks(client *httpClient.HttpClient, externalLinks []string) []string {
	var wg sync.WaitGroup
	var response []string
	output := make(chan string)
	for _, url := range externalLinks {
		wg.Add(1)
		go checkURL(client, url, &output, &wg)
	}
	go func() {
		wg.Wait()
		close(output)
	}()
	for url := range output {
		if url != "" {
			fmt.Println(url)
			response = append(response, url)
		}
	}
	return response
}

//Create a Get request for each URL and check http status without processing response
func checkURL(client *httpClient.HttpClient, url string, ch *chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	status, err := client.MakeRequestWithoutResponseBody(url)
	if err != nil {
		*ch <- url
		return
	}
	if status == http.StatusOK {
		*ch <- ""
	} else {
		*ch <- url
	}
}

//SliceContains returns true if `slice` contains `value`
func SliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

//Parse the HTML node tree and find headings,Form and Links
func ParseHtml(n *html.Node, headings map[string]int, links []string, isForm *bool) []string {

	if _, exists := headings[n.Data]; exists {
		headings[n.Data]++
	}
	if n.Data == "form" {
		*isForm = true
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				if !SliceContains(links, a.Val) {
					links = append(links, a.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = ParseHtml(c, headings, links, isForm)
	}
	return links
}
