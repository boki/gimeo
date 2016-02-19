package gimeo

import (
	"fmt"

	. "github.com/julianedialkova/gimeo/data"
)

func (c *Client) GetGroups(params *Parameters) (*GroupData, error) {
	resp, err := c.Get("/groups", params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &GroupData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

////
func (c *Client) CreateGroup(params *Parameters) (*GroupDataElement, error) {
	resp, err := c.Post("/groups", params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &GroupDataElement{}

	err = c.processRequestData(204, resp, data)
	return data, err
}

func (c *Client) GetGroup(group string, params *Parameters) (*GroupData, error) {
	uri := fmt.Sprintf("/groups/%s", group)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &GroupData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) DeleteGroup(group string, params *Parameters) error {
	uri := fmt.Sprintf("/groups/%s", group)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

func (c *Client) GetGroupUsers(group string, params *Parameters) (*UserData, error) {
	uri := fmt.Sprintf("/groups/%s/users", group)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &UserData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) GetGroupVideos(group string, params *Parameters) (*VideoData, error) {
	uri := fmt.Sprintf("/groups/%s/videos", group)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &VideoData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) CheckIfVideoInGroup(group, video string, params *Parameters) (*VideoDataElement, error) {
	uri := fmt.Sprintf("/groups/%s/videos/%s", group, video)
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
func (c *Client) PutVideoInGroup(group, video string, params *Parameters) error {
	uri := fmt.Sprintf("/groups/%s/videos/%s", group, video)
	resp, err := c.Put(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

func (c *Client) DeleteVideoFromGroup(group, video string, params *Parameters) error {
	uri := fmt.Sprintf("/groups/%s/videos/%s", group, video)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}
	err = c.processRequestNoData(resp)
	return err
}
