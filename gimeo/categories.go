package gimeo

import (
	"fmt"

	. "github.com/julianedialkova/gimeo/data"
)

//Get a list of the top level categories.
func (c *Client) GetCategories(params *Parameters) (*CategoryData, error) {
	resp, err := c.Get("/categories", params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &CategoryData{}
	err = c.processRequestData(200, resp, data)
	return data, err
}

//Get a category.
func (c *Client) GetCategory(category string, params *Parameters) (*CategoryDataElement, error) {
	uri := fmt.Sprintf("/categories/%s", category)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &CategoryDataElement{}
	err = c.processRequestData(200, resp, data)
	return data, err
}

// Get a list of Channels related to a category.
func (c *Client) GetCategoryChannels(category string, params *Parameters) (*ChannelData, error) {
	uri := fmt.Sprintf("/categories/%s/channels", category)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &ChannelData{}
	err = c.processRequestData(200, resp, data)
	return data, err
}

// Get a list of Groups related to a category.
func (c *Client) GetCategoryGroups(group string, params *Parameters) (*GroupData, error) {
	uri := fmt.Sprintf("/categories/%s/groups", group)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &GroupData{}
	err = c.processRequestData(200, resp, data)
	return data, err
}

// Get a list of videos related to a category.
func (c *Client) GetCategoryVideos(category string, params *Parameters) (*VideoData, error) {
	uri := fmt.Sprintf("/categories/%s/videos", category)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}
	data := &VideoData{}
	err = c.processRequestData(200, resp, data)
	return data, err
}

// Check if a category contains a video
func (c *Client) CheckIfVideoInCategory(category, video string, params *Parameters) (*VideoDataElement, error) {
	uri := fmt.Sprintf("/categories/%s/videos/%s", category, video)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &VideoDataElement{}
	err = c.processRequestData(200, resp, data)
	return data, err
}
