package controllers

import (
	"fmt"
	"htmlparser/entities/model"
	"htmlparser/infrastructure/httpClient"
	"htmlparser/parser"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	client httpClient.HttpClient
	parse  parser.Parser
)

func Initialize() {
	client.Initialize()
	var htmlParser parser.HtmlParser
	htmlParser.Client = &client
	parse = htmlParser
}

func ParseHandler(c *gin.Context) {
	var url model.URL
	c.BindJSON(&url)
	request, err := client.CreateRequest(url.URL, "GET")
	if err != nil {
		c.AbortWithError(500, fmt.Errorf("unable to create HTTP request"))
	}
	reader, err := client.MakeRequest(request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("request failed"))
	}

	response, err := parse.ProcessHTML(reader)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("request failed"))
	}

	c.JSON(http.StatusOK, response)
}
