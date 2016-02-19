package gimeo

import (
	"fmt"

	. "github.com/julianedialkova/gimeo/data"
)

func (c *Client) GetTag(tag string, params *Parameters) (*Tag, error) {
	uri := fmt.Sprintf("/tags/%s", tag)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &Tag{}

	err = c.processRequestData(200,resp, data)
	return data, err
}

func (c *Client) GetTagVideos(tag string, params *Parameters) (*VideoData, error) {
	uri := fmt.Sprintf("/tags/%s/videos", tag)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &VideoData{}

	err = c.processRequestData(200,resp, data)
	return data, err
}
