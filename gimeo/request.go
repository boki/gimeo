package gimeo

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Parameters map[string]string

func (params Parameters) ToUrlValues() url.Values {
	result := url.Values{}
	for key, value := range params {
		result.Set(key, value)
	}
	return result
}

type Headers map[string]string

type Request struct {
	Hostname string
	Headers  *Headers
}

var DefaultRequest = &Request{
	Hostname: "https://api.vimeo.com",
	Headers: &Headers{
		"Content-Type": "application/x-www-form-urlencoded",
		"Accept":       "application/vnd.vimeo.*+json;version=3.2",
	},
}

func buildRequestUrl(urlStr string) string {
	return fmt.Sprintf("%s/%s", DefaultRequest.Hostname, strings.TrimPrefix(urlStr, "/"))
}

func performRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Do(req)
}

func (c *Client) applyDefaults(req *http.Request) {
	for key, value := range *(DefaultRequest.Headers) {
		req.Header.Set(key, value)
	}
	if c.accessToken == "" {
		auth := []byte(fmt.Sprintf("%s:%s", c.clientID, c.clientSecret))
		authHeader := base64.StdEncoding.EncodeToString(auth)
		req.Header.Set("Authorization", fmt.Sprintf("basic %s", authHeader))
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("bearer %s", c.accessToken))
}

func (c *Client) Get(urlStr string, params *Parameters) (*http.Response, error) {

	requestUri := buildRequestUrl(urlStr)
	req, err := http.NewRequest("GET", requestUri, nil)

	if err != nil {
		return nil, err
	}

	c.applyDefaults(req)
	fmt.Printf("Performing GET request to %s\n", requestUri)
	if params != nil {
		encoded := params.ToUrlValues()
		req.URL.RawQuery = encoded.Encode()
	}

	return performRequest(req)

}

func (c *Client) Patch(urlStr string, params *Parameters) (*http.Response, error) {

	requestUri := buildRequestUrl(urlStr)

	encoded := params.ToUrlValues()
	body := encoded.Encode()

	req, err := http.NewRequest("PATCH", requestUri, strings.NewReader(body))

	if err != nil {
		return nil, err
	}
	c.applyDefaults(req)

	fmt.Printf("Performing PATCH request to %s\n", requestUri)

	return performRequest(req)
}

func (c *Client) Post(urlStr string, params *Parameters) (*http.Response, error) {

	requestUri := buildRequestUrl(urlStr)

	encoded := params.ToUrlValues()
	body := encoded.Encode()

	req, err := http.NewRequest("POST", requestUri, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	c.applyDefaults(req)

	fmt.Printf("Performing POST request to %s\n", requestUri)

	return performRequest(req)

}

func (c *Client) Put(urlStr string) (*http.Response, error) {

	requestUri := buildRequestUrl(urlStr)
	req, err := http.NewRequest("PUT", requestUri, nil)

	if err != nil {
		return nil, err
	}

	c.applyDefaults(req)

	fmt.Printf("Performing PUT request to %s\n", requestUri)

	return performRequest(req)

}

func (c *Client) Delete(urlStr string) (*http.Response, error) {

	requestUri := buildRequestUrl(urlStr)
	req, err := http.NewRequest("DELETE", requestUri, nil)

	if err != nil {
		return nil, err
	}

	c.applyDefaults(req)

	fmt.Printf("Performing DELETE request to %s\n", requestUri)

	return performRequest(req)
}

func (c *Client) processRequestData(status int, resp *http.Response, data interface{}) error {
	if resp.StatusCode != status {
		fmt.Printf("Request returned %s.\n", resp.Status)
		return fmt.Errorf("Request returned %s.\n", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Could not read requests body")
		return err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Could not parse body")
		return err
	}

	return nil
}

func (c *Client) processRequestNoData(resp *http.Response) error {
	if resp.StatusCode != 204 {
		fmt.Printf("Request returned %s.\n", resp.Status)
		return fmt.Errorf("Request returned %s.\n", resp.Status)
	}
	return nil
}
