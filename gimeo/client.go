package gimeo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
)

const (
	libVersion = "v1.0.0"
	userAgent  = "gimeo/" + libVersion + " (" + runtime.GOOS + "/" + runtime.GOARCH + ")"
)

//Logger is the interface for log operations.
type Logger interface {
	//Printf prints to the logger. Arguments are handled in the manner of fmt.Printf.
	Printf(format string, v ...interface{})
	//Println prints to the logger. Arguments are handled in the manner of fmt.Println.
	Println(v ...interface{})
}

//Client is the type of object that identifies the user.
type Client struct {
	UserAgent    string // optional additional User-Agent fragment
	clientID     string
	clientSecret string
	accessToken  string
	log          Logger
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
func Vimeo(clientID, clientSecret, accessToken string, logger ...Logger) *Client {
	if len(logger) == 0 || logger[0] == nil {
		logger = []Logger{log.New(os.Stderr, "gimeo: ", log.LstdFlags)}
	}
	return &Client{
		clientID:     clientID,
		clientSecret: clientSecret,
		accessToken:  accessToken,
		log:          logger[0],
	}
}

// UserAgent returns the user agent to use for requests.
func (c *Client) userAgent() string {
	if c.UserAgent == "" {
		return userAgent
	}
	return c.UserAgent + " " + userAgent
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
		response := make(map[string]interface{})
		err = json.Unmarshal(body, &response)
		if err != nil {
			return err
		}
		c.accessToken = response["access_token"].(string)
	}
	return nil
}
