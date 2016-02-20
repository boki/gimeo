package gimeo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Client is the type of object that identifies the user.
type Client struct {
	clientID     string
	clientSecret string
	accessToken  string
}

//AuthEndpoints are the main authorization endpoints
type AuthEndpoints struct {
	athorization      string
	accessToken       string
	clientCredentials string
}

// AuthenticationEndpoints contains the default auth endpoints
var AuthenticationEndpoints = &AuthEndpoints{
	"/oauth/authorize",
	"/oauth/access_token",
	"/oauth/authorize/client",
}

// Vimeo creates a client to interact with the Vimeo API
func Vimeo(clientID, clientSecret, accessToken string) *Client {
	return &Client{
		clientID:     clientID,
		clientSecret: clientSecret,
		accessToken:  accessToken,
	}
}

// GenerateAuthAccessToken is the first step of the redirect process is to send the user's client (browser) to vimeo.com. This is generally accomplished by providing the authorize url as a link on a webpage.
func (c *Client) GenerateAuthAccessToken(redirectURI, scope, state string) error {

	params := &Parameters{
		"response_type": "code",
		"client_id":     c.clientID,
		"scope":         scope,
		"redirect_uri":  redirectURI,
		"state":         state,
	}

	_, err := c.Post(AuthenticationEndpoints.athorization, params)

	return err
}

//GetToken exchange a code for an access token. This code should exist on your redirect_uri
func (c *Client) GetToken(code, redirectURI string) error {

	params := &Parameters{
		"grant_type":   "authorization_code",
		"code":         code,
		"redirect_uri": redirectURI,
	}

	resp, err := c.Post(AuthenticationEndpoints.accessToken, params)
	if err != nil {
		return err
	}
	return c.setToken(resp)
}

// GenerateUnauthAccessToken generates an access token it order to make authorized requests. The function takes code and a redirect_uri and returns an access_token.
func (c *Client) GenerateUnauthAccessToken(scope string) error {

	params := &Parameters{
		"grant_type": "client_credentials",
		"scope":      scope,
	}
	resp, err := c.Post(AuthenticationEndpoints.clientCredentials, params)

	if err != nil {
		return err
	}
	return c.setToken(resp)
}

func (c *Client) setToken(resp *http.Response) error {
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		response := make(map[string]string)
		err = json.Unmarshal(body, &response)
		if err != nil {
			return err
		}
		c.accessToken = response["access_token"]
		fmt.Println(response)
	}

	return nil
}
