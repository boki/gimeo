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

	// fi, err := file.Stat()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// req.Header.Set("Content-Length", string(fi.Size()))
	req.Header.Set("Content-Type", "video/mp4")

	resp, err := performRequest(req)

	c.Verify(structer.CompleteURI)
	// fmt.Println(resp.StatusCode)
	// fmt.Println(resp.Status)
	return resp, err
}

func (c *Client) VerifyUpload(uri string) (*http.Response, error) {
	// uri := structer.UploadLink
	// uri := "http://1511923893.cloud.vimeo.com/upload?ticket_id=96b6144ad2ec49f51f97e83c2489be34&video_file_id=483785657&signature=23fa1e0ae87f9f9b14db6ea4ca42ea5c&v6=1"
	// uri := fmt.Sprintf("https://1234.cloud.vimeo.com/upload?ticket_id=%s", structer.TicketID)
	req, err := http.NewRequest("PUT", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("bearer %s", c.accessToken))
	req.Header.Set("Content-Range", "bytes */*")
	req.Header.Set("Content-Length", string(0))
	req.Header.Set("Content-Type", "video/mp4")

	resp, err := performRequest(req)
	fmt.Println(resp)

	bod, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Could not read requests body")
		return nil, err
	}

	fmt.Println(bod)
	return resp, err
}

func (c *Client) Verify(completeUri string) (*http.Response, error) {
	resp, err := c.Delete(completeUri)
	bod, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Could not read requests body")
		return nil, err
	}

	fmt.Println(string(bod))
	fmt.Println(resp.Status)

	return resp, nil
}
