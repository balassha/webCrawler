package httpClient

import (
	"crypto/tls"
	"io"
	"net/http"
	"time"
)

type HttpClient struct {
	Client *http.Client
}

func (c *HttpClient) Initialize() {

	// To handle websites with Self-Signed certificates
	// Can be ignored if this support is not required
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c.Client = &http.Client{
		Timeout:   10 * time.Second,
		Transport: tr,
	}
}

// Provides the ability to customize the Request
func (c *HttpClient) CreateRequest(url string, method string) (*http.Request, error) {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	//request.Header.Set("User-Agent", "Chrome")
	return request, nil
}

func (c *HttpClient) MakeRequest(request *http.Request) (io.Reader, error) {
	response, err := c.Client.Do(request)
	if err != nil {
		return nil, err
	}

	return response.Body, nil
}

func (c *HttpClient) MakeRequestWithoutResponseBody(url string) (int, error) {
	request, err := c.CreateRequest(url, "GET")
	if err != nil {
		return 500, err
	}
	response, err := c.Client.Do(request)
	if err != nil {
		return 500, err
	}
	return response.StatusCode, nil
}
