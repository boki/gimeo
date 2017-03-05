package gimeo

import (
	"fmt"

	. "github.com/julianedialkova/gimeo/data"
)

//GetGroups gets a list of all groups
// Name								Type				Required			Description
// page								int					No						The page number to show.
// per_page						int					No						Number of items to show on each page. Max 50.
// query							string			No						Search query.
// sort								string			No						Technique used to sort the results.
// 										date
// 										alphabetical
// 										videos
// 										followers
// direction					string			No						The direction that the results are sorted.
// 										asc
// 										desc
// filter							string 			No 						Filter to apply to the results.
//										featured
func (c *Client) GetGroups(params *Parameters) (*GroupData, error) {
	resp, err := c.Get("/groups", params)

	if err != nil {
		return nil, err
	}

	data := &GroupData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// CreateGroup creates a new group
// This method requires a token with the "create" scope.
func (c *Client) CreateGroup(params *Parameters) (*GroupDataElement, error) {
	resp, err := c.Post("/groups", params)

	if err != nil {
		return nil, err
	}

	data := &GroupDataElement{}

	err = c.processRequestData(204, resp, data)
	return data, err
}

// GetGroup gets a group
func (c *Client) GetGroup(group string, params *Parameters) (*GroupData, error) {
	uri := fmt.Sprintf("/groups/%s", group)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &GroupData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// DeleteGroup deletes a GetGroup
// This method requires a token with the "delete" scope.
func (c *Client) DeleteGroup(group string, params *Parameters) error {
	uri := fmt.Sprintf("/groups/%s", group)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

// GetGroupUsers gets a list of users that joined a Group.
// Name				Type					Required	Description
// page				int						No				The page number to show.
// per_page		int						No				Number of items to show on each page. Max 50.
// query			string				No				Search query.
// sort				string				No				Technique used to sort the results.
//						date
// 						alphabetical
// direction	string				No				The direction that the results are sorted.
// 						asc
// 						desc
// filter			string				No				Filter to apply to the results.
// 						moderators
func (c *Client) GetGroupUsers(group string, params *Parameters) (*UserData, error) {
	uri := fmt.Sprintf("/groups/%s/users", group)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &UserData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

//GetGroupVideos gets a list of videos in a Group.
// Name								Type				Required			Description
// page								int					No						The page number to show.
// per_page						int					No						Number of items to show on each page. Max 50.
// query							string			No						Search query.
// filter							string			No						Filter to apply to the results.
// 										embeddable
// filter_embeddable	string			No						Required if filter=embeddable. Choose between only videos that are embeddable, and only videos that are not embeddable.
// 										true
// 										false
// sort								string			No						Technique used to sort the results.
// 										date
// 										alphabetical
// 										plays
// 										likes
// 										comments
// 										duration
// direction					string			No						The direction that the results are sorted.
// 										asc
// 										desc
func (c *Client) GetGroupVideos(group string, params *Parameters) (*VideoData, error) {
	uri := fmt.Sprintf("/groups/%s/videos", group)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &VideoData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// CheckIfVideoInGroup checks if a Group has a video.
func (c *Client) CheckIfVideoInGroup(group, video string, params *Parameters) (*VideoDataElement, error) {
	uri := fmt.Sprintf("/groups/%s/videos/%s", group, video)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &VideoDataElement{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// PutVideoInGroup adds a video to a Group
// This method requires a token with the "edit" scope.
func (c *Client) PutVideoInGroup(group, video string, params *Parameters) error {
	uri := fmt.Sprintf("/groups/%s/videos/%s", group, video)
	resp, err := c.Put(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

// DeleteVideoFromGroup removes a video from a Group
// This method requires a token with the "edit" scope.
func (c *Client) DeleteVideoFromGroup(group, video string, params *Parameters) error {
	uri := fmt.Sprintf("/groups/%s/videos/%s", group, video)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}
	err = c.processRequestNoData(resp)
	return err
}
