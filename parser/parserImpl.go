package parser

import (
	"bytes"
	"htmlparser/infrastructure/httpClient"
	"htmlparser/webscrapper"
	"io"
	"io/ioutil"

	"golang.org/x/net/html"
)

type HtmlParser struct {
	Client *httpClient.HttpClient
}

func (h HtmlParser) ProcessHTML(data io.Reader) (map[string]interface{}, error) {
	//Create HTML Page String and HTML Tree of html.Nodes
	buf := &bytes.Buffer{}
	tee := io.TeeReader(data, buf)

	node, err := html.Parse(tee)
	if err != nil {
		return nil, err
	}

	//Get the response body as a string
	dataInBytes, err := ioutil.ReadAll(buf)
	if err != nil {
		return nil, err
	}

	// Initialize webscrapper library
	var scrapper webscrapper.Scrapper
	scrapper.Node = node
	scrapper.HtmlPage = string(dataInBytes)

	//Create empty response
	response := make(map[string]interface{})

	//Get HTML Version
	response["version"], err = scrapper.GetHTMLVersion()
	if err != nil {
		return nil, err
	}

	//Get Title
	response["title"], err = scrapper.GetTitle()
	if err != nil {
		return nil, err
	}

	//Initiate HTML Parsing to get further data
	scrapper.ParseHtmlFile()

	//Get Headings count
	response["headingsCount"], err = scrapper.GetHeadingsCount()
	if err != nil {
		return nil, err
	}

	//Get Internal and External Links
	response["links"], err = scrapper.GetInternalAndExternalLinks()
	if err != nil {
		return nil, err
	}

	//Get Inaccessible Links
	response["inaccessibleLinks"], err = scrapper.GetInaccessibleLinks(h.Client)
	if err != nil {
		return nil, err
	}

	//Check if the html has a Form
	response["isLoginForm"], err = scrapper.IsLoginPage()
	if err != nil {
		return nil, err
	}

	return response, nil
}
