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

//Parameters given to the functions
type Parameters map[string]string

// BuildRequestURL parses Parameters to url.Values
func (params Parameters) BuildRequestURL() url.Values {
	result := url.Values{}
	for key, value := range params {
		result.Set(key, value)
	}
	return result
}

// Headers structure
type Headers map[string]string

//Request structure
type Request struct {
	Hostname string
	Headers  *Headers
}

// DefaultRequest contains the default values, used by all requests
var DefaultRequest = &Request{
	Hostname: "https://api.vimeo.com",
	Headers: &Headers{
		"Content-Type": "application/x-www-form-urlencoded",
		"Accept":       "application/vnd.vimeo.*+json;version=3.2",
	},
}

//BuildRequestURL builds a request URL from host and path
func BuildRequestURL(urlStr string) string {
	return fmt.Sprintf("%s/%s", DefaultRequest.Hostname, strings.TrimPrefix(urlStr, "/"))
}

func performRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Do(req)
}

//ApplyDefaults applies the default values from DefaultRequest to our request
func (c *Client) ApplyDefaults(req *http.Request) {
	for key, value := range *(DefaultRequest.Headers) {
		req.Header.Set(key, value)
	}
	req.Header.Set("User-Agent", c.userAgent())
	if c.accessToken == "" {
		auth := []byte(fmt.Sprintf("%s:%s", c.clientID, c.clientSecret))
		authHeader := base64.StdEncoding.EncodeToString(auth)
		req.Header.Set("Authorization", fmt.Sprintf("basic %s", authHeader))
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("bearer %s", c.accessToken))
}

// Get performs a GET request to the Vimeo API with urlStr path and params as body.
func (c *Client) Get(urlStr string, params *Parameters) (*http.Response, error) {

	requestURI := BuildRequestURL(urlStr)
	req, err := http.NewRequest("GET", requestURI, nil)

	if err != nil {
		return nil, err
	}

	c.ApplyDefaults(req)
	c.log.Printf("Performing GET request to %s\n", requestURI)
	if params != nil {
		encoded := params.BuildRequestURL()
		req.URL.RawQuery = encoded.Encode()
	}

	return performRequest(req)

}

// Patch performs a PATCH request to the Vimeo API with urlStr path and params as body.
func (c *Client) Patch(urlStr string, params *Parameters) (*http.Response, error) {

	requestURI := BuildRequestURL(urlStr)

	encoded := params.BuildRequestURL()
	body := encoded.Encode()

	req, err := http.NewRequest("PATCH", requestURI, strings.NewReader(body))

	if err != nil {
		return nil, err
	}
	c.ApplyDefaults(req)

	c.log.Printf("Performing PATCH request to %s\n", requestURI)

	return performRequest(req)
}

// Post performs a POST request to the Vimeo API with urlStr path and params as body.
func (c *Client) Post(urlStr string, params *Parameters) (*http.Response, error) {

	requestURI := BuildRequestURL(urlStr)

	encoded := params.BuildRequestURL()
	body := encoded.Encode()

	req, err := http.NewRequest("POST", requestURI, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	c.ApplyDefaults(req)

	c.log.Printf("Performing POST request to %s\n", requestURI)

	return performRequest(req)

}

// Put performs a PUT request to the Vimeo API with urlStr path and params as body.
func (c *Client) Put(urlStr string) (*http.Response, error) {

	requestURI := BuildRequestURL(urlStr)
	req, err := http.NewRequest("PUT", requestURI, nil)

	if err != nil {
		return nil, err
	}

	c.ApplyDefaults(req)

	c.log.Printf("Performing PUT request to %s\n", requestURI)

	return performRequest(req)

}

// Delete performs a DELETE request to the Vimeo API with urlStr path and params as body.
func (c *Client) Delete(urlStr string) (*http.Response, error) {

	requestURI := BuildRequestURL(urlStr)
	req, err := http.NewRequest("DELETE", requestURI, nil)

	if err != nil {
		return nil, err
	}

	c.ApplyDefaults(req)

	c.log.Printf("Performing DELETE request to %s\n", requestURI)

	return performRequest(req)
}

func (c *Client) processRequestData(status int, resp *http.Response, data interface{}) error {
	if resp.StatusCode != status {
		c.log.Printf("Request returned %s.\n", resp.Status)
		return fmt.Errorf("request returned %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) processRequestNoData(resp *http.Response) error {
	if resp.StatusCode != 204 {
		c.log.Printf("Request returned %s.\n", resp.Status)
		return fmt.Errorf("request returned %s", resp.Status)
	}
	return nil
}
