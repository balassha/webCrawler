package webscrapper

import (
	"fmt"
	"htmlparser/entities/model"
	"htmlparser/infrastructure/httpClient"
	"strings"

	"golang.org/x/net/html"
)

type Scrapper struct {
	HtmlPage          string
	Node              *html.Node
	Headings          map[string]int
	Links             []string
	IsLogin           bool
	ExternalLinks     []string
	InaccessibleLinks []string
}

func (s *Scrapper) UpdateHeadings() {
	s.Headings = map[string]int{
		"h1": 0,
		"h2": 0,
		"h3": 0,
		"h4": 0,
		"h5": 0,
		"h6": 0,
	}
}

//Gets HTML version from page
func (s *Scrapper) GetHTMLVersion() (string, error) {
	version := GetHTMLVersion(s.HtmlPage)
	if version == "UNKNOWN" {
		return "", fmt.Errorf("unknown HTML Version")
	}
	return version, nil
}

//Gets HTML title from page
func (s *Scrapper) GetTitle() (string, error) {
	return GetPageTitle(s.Node), nil
}

//Return the count of headings
func (s *Scrapper) GetHeadingsCount() (map[string]int, error) {
	return s.Headings, nil
}

//Parses the HTML node tree and gets relevant information
func (s *Scrapper) ParseHtmlFile() {
	s.Links = ParseHtml(s.Node, s.Headings, nil, &s.IsLogin)
}

//Identify Internal and External links
func (s *Scrapper) GetInternalAndExternalLinks() (model.Links, error) {
	var response model.Links
	if s.Links == nil {
		return response, fmt.Errorf("links not initiaized")
	}
	for _, value := range s.Links {
		if strings.Contains(value, "http") {
			response.ExternalLinks = append(response.ExternalLinks, value)
			s.ExternalLinks = append(s.ExternalLinks, value)
		} else {
			response.InternalLinks = append(response.InternalLinks, value)
		}
	}
	return response, nil
}

//Return whether the html page has a login form
func (s *Scrapper) IsLoginPage() (bool, error) {
	return s.IsLogin, nil
}

//Gets all Inaccessible external links in the html page
func (s *Scrapper) GetInaccessibleLinks(client *httpClient.HttpClient) ([]string, error) {
	s.InaccessibleLinks = ProcessInaccessibleLinks(client, s.ExternalLinks)
	fmt.Println("Inaccess", s.InaccessibleLinks)
	return s.InaccessibleLinks, nil
}
