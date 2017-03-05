package gimeo

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

//Upload uploads a video to Vimeo
func (c *Client) Upload(filePath string) (*http.Response, error) {
	params := &Parameters{
		"type": "streaming",
	}
	structer, err := c.MeUploadVideo(params)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	//
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", structer.UploadLinkSecure, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("bearer %s", c.accessToken))
	req.Header.Set("Content-Type", "video/mp4")

	resp, err := performRequest(req)

	c.Verify(structer.CompleteURI)
	return resp, err
}

//Verify verifies that the upload was successful
func (c *Client) Verify(completeURI string) (*http.Response, error) {
	resp, err := c.Delete(completeURI)
	bod, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(bod))
	fmt.Println(resp.Status)

	return resp, nil
}
