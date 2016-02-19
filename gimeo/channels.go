package gimeo

import (
	"fmt"

	. "github.com/julianedialkova/gimeo/data"
)

func (c *Client) GetChannels(params *Parameters) (*ChannelData, error) {
	resp, err := c.Get("/channels", params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &ChannelData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// Do not have permissions ?
/////
func (c *Client) CreateChannel(params *Parameters) (map[string]interface{}, error) {
	resp, err := c.Post("/channels", params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := make((map[string]interface{}))

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) GetChannel(channel string, params *Parameters) (*ChannelDataElement, error) {
	uri := fmt.Sprintf("/channels/%s", channel)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &ChannelDataElement{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

/////
func (c *Client) PatchChannel(channel string, params *Parameters) (*ChannelDataElement, error) {
	uri := fmt.Sprintf("/channels/%s", channel)
	resp, err := c.Patch(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &ChannelDataElement{}

	err = c.processRequestData(204, resp, data)
	return data, err
}

/////
func (c *Client) DeleteChannel(channel string, params *Parameters) (*ChannelDataElement, error) {
	uri := fmt.Sprintf("/channels/%s", channel)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &ChannelDataElement{}

	err = c.processRequestData(204, resp, data)
	return data, err
}

//////
func (c *Client) GetChannelUsers(channel string, params *Parameters) (map[string]interface{}, error) {
	uri := fmt.Sprintf("/channels/%s/users", channel)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := make(map[string]interface{})

	err = c.processRequestData(200, resp, data)
	return data, err
}

// GEThttps://api.vimeo.com/channels/{channel_id}/videos
func (c *Client) GetChannelVideos(channel string, params *Parameters) (*VideoData, error) {
	uri := fmt.Sprintf("/channels/%s/videos", channel)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &VideoData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) CheckIfVideoInChannel(channel, video string, params *Parameters) (*VideoDataElement, error) {
	uri := fmt.Sprintf("/channels/%s/videos/%s", channel, video)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &VideoDataElement{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

///////
func (c *Client) PutVideoInChannel(channel, video string, params *Parameters) error {
	uri := fmt.Sprintf("/channels/%s/videos/%s", channel, video)
	resp, err := c.Put(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

//////
func (c *Client) DeleteVideoFromChannel(channel, video string, params *Parameters) error {
	uri := fmt.Sprintf("/channels/%s/videos/%s", channel, video)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}
