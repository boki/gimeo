package gimeo

import (
	// . "data"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	// . "github.com/julianedialkova/gimeo/data"
	// "encoding/base64"
	// "net/url"
	// "strings"
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

var AuthenticationEndpoints = &AuthEndpoints{
	"/oauth/authorize",
	"/oauth/access_token",
	"/oauth/authorize/client",
}

/*
Vimeo - used to interact with the Vimeo API
*/
func Vimeo(clientID, clientSecret, accessToken string) *Client {
	return &Client{
		clientID:     clientID,
		clientSecret: clientSecret,
		accessToken:  accessToken,
	}
}

//Exchange a code for an access token. This code should exist on your redirect_uri

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

func (c *Client) getToken(code, redirectURI string) error {

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

/*
 * The first step of the authorization process.
 *
 * This function returns a url, which the user should be sent to (via redirect or link).
 * The destination allows the user to accept or deny connecting with vimeo, and accept or deny each of the scopes you requested.
 * Scopes are passed through the second parameter as an array of strings, or a space delimited list.
 *
 * Once accepted or denied, the user is redirected back to the redirect_uri.
 * If accepted, the redirect url will
 */
func (c *Client) GenerateAuthAccessToken(redirectURI, scope, state string) error {

	params := &Parameters{
		"response_type": "code",
		"client_id":     c.clientID,
		"scope":         scope,
		"redirect_uri":  redirectURI,
		"state":         state,
	}

	_, err := c.Post(AuthenticationEndpoints.athorization, params)

	if err != nil {
		return err
	}

	return nil

}

/*
GenerateUnauthAccessToken generates an access token it order to make authorized requests
The function takes code and a redirect_uri and returns an access_token.
*/

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
