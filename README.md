# gimeo
This is a Go library to consume the Vimeo API. More information on the API can be found on the [official page](https://developer.vimeo.com/api/start).

## Installation
All you need to do is run:
```
go get https://github.com/julianedialkova/gimeo.git
```
Then you can enjoy the library!

## Using the library

### First steps
In order to make requests to the API you must first run the following function:

```
var client = Vimeo("clientID", "clientSecret", "accessToken")
```

You should provide your client ID and secret. They can be found on your app page under the Authentication tab. You can create an app [here](https://developer.vimeo.com/apps).

You can also provide an access token, but it is optional. It is required in order to make requests. Below you can find information how to generate one.

### Generate an access token

There are two ways to make requests to the Vimeo API - through authenticated and unauthenticated tokens. First you should have called the Vimeo function to generate a client.

1. Unauthenticated token

  This is a token without a user. All you can see using this token is public data. If you want to generate one, just use:

  ```
  client.GenerateUnauthAccessToken([]string{"create"})
  ```
  Where scopes look like this: `[]string{"public"}`.

  As an argument you should pass the scopes you want to access. Description of the scopes can be found on the [official page](https://developer.vimeo.com/api/authentication#scopes).

1. Authenticated token

  This token uses the user. It interacts on behalf of the authenticated user. In order to generate one, you should:

  ```
  client.GenerateАuthAccessToken(scopes, "redirectURI", "state")
  ```
  Where scopes look like this: `[]string{"public"}`.

|  Name         |  Description
|--------------|-------------
|  redirectURI  | This must be required, and must match your app callback URL
|  state        | A unique value which the client will return alongside access tokens



### Make a request
When you have all ready making a request is really simple:

```
client.Request(`
  {
    "path"  : "/me/videos",
    "query" : {
        "per_page" : "1"
  }
`)
```

Values you can put in the request are: `method`, `path`, `query` and `headers`.


### Upload a video
You may also upload a video or replace it.

1. Upload a video
  ```
  client.uploadVideo("/path/to/video.mp4", nil)
  ```

1. Replace a video
  ```
  client.uploadVideo("/path/to/video.mp4", "/videos/videoToReplace")
  ```
