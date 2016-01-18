package vimeo

import "net/http"

//AuthEndpoints are the main authorization endpoints
/* Default ones will be:
// 	authorization : '/oauth/authorize',
// 	accessToken : '/oauth/access_token',
// 	clientCredentials : '/oauth/authorize/client'
*/
type AuthEndpoints struct {
	athorization      string
	accessToken       string
	clientCredentials string
}

//Client is the type of object that identifies the user.
type Client struct {
	clientID     string
	clientSecret string
	accessToken  string
}

/*
Request shows the structure of a request
Default ones will be:
protocol : 'https:',
hostname : 'api.vimeo.com',
port : 443,
method : 'GET',
query : {},
headers : {
	Accept: "application/vnd.vimeo.*+json;version=3.2",
	'User-Agent': 'Vimeo.js/1.1.4'
*/
type Request struct {
	protocol string
	hostname string
	Path     string `json:"path"`
	port     int
	Method   string `json:"method"`
	Query    map[string]string
	Headers  map[string]string
}

/*
Vimeo - used to interact with the Vimeo API
*/
func Vimeo(clientID, clientSecret, accessToken string) *Client {}

/*
Request performs an api call. There are two ways for that:

1. Provide only URL
	If an URL is provided, then the default settings will be set. => GET https://api.vimeo.com/{url provided}

2. Options
	If there is a options map provided, then those elements will be added to the request query.
	The only required one is options["path"].Options can include hostname, port,
	query (applied to the URL if the method is GET and in the request body if POST), headers,
	path (can also include a querystring) and method.

The function will return the response body and error.
*/
func (c *Client) Request(options []byte) (response []byte, err error) {}

/*
buildRequest builds the actual request after merging the provided options with the default ones
and returns a Request objectfrom the options provided
*/
func (c *Client) buildRequest(options map[string]string) *http.Request {}

/*
GenerateAuthAccessToken generates an access token it order to make authorized requests
The function takes code and a redirect_uri and returns an access_token.
*/
func (c *Client) GenerateAuthAccessToken(code, redirectURI, state string) string {}

/*
GenerateUnauthAccessToken generates an access token it order to make authorized requests
The function takes code and a redirect_uri and returns an access_token.
*/
func (c *Client) GenerateUnauthAccessToken(scopes []string) {}

/*
UploadVideo is used in order to upload a video to Vimeo
*/
func (c *Client) UploadVideo(path, videoURI string) {}
