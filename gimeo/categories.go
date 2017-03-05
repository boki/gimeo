package gimeo

import (
	"fmt"

	. "github.com/julianedialkova/gimeo/data"
)

//GetCategories gets a list of the top level categories.
// Possible parameters
// Name			Type	Required	Description
// page			int		No				The page number to show.
// per_page	int		No				Number of items to show on each page. Max 50.
func (c *Client) GetCategories(params *Parameters) (*CategoryData, error) {
	resp, err := c.Get("/categories", params)

	if err != nil {
		return nil, err
	}

	data := &CategoryData{}
	err = c.processRequestData(200, resp, data)
	return data, err
}

// GetCategory gets a category.
func (c *Client) GetCategory(category string, params *Parameters) (*CategoryDataElement, error) {
	uri := fmt.Sprintf("/categories/%s", category)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &CategoryDataElement{}
	err = c.processRequestData(200, resp, data)
	return data, err
}

// GetCategoryChannels gets a list of Channels related to a category.
// Name				Type					Required	Description
// page				int						No				The page number to show.
// per_page		int						No				Number of items to show on each page. Max 50.
// query			string				No				Search query.
// sort				string				Но				No	Technique used to sort the results.
//						date
// 						alphabetical
// 						videos
// 						followers
// direction	string				Но				No	The direction that the results are sorted.
// 						asc
// 						desc
func (c *Client) GetCategoryChannels(category string, params *Parameters) (*ChannelData, error) {
	uri := fmt.Sprintf("/categories/%s/channels", category)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &ChannelData{}
	err = c.processRequestData(200, resp, data)
	return data, err
}

// GetCategoryGroups get a list of Groups related to a category.
// Name				Type					Required	Description
// page				int						No				The page number to show.
// per_page		int						No				Number of items to show on each page. Max 50.
// query			string				No				Search query.
// sort				string				Но				No	Technique used to sort the results.
//						date
// 						alphabetical
// 						videos
// 						followers
// direction	string				Но				No	The direction that the results are sorted.
// 						asc
// 						desc
func (c *Client) GetCategoryGroups(group string, params *Parameters) (*GroupData, error) {
	uri := fmt.Sprintf("/categories/%s/groups", group)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &GroupData{}
	err = c.processRequestData(200, resp, data)
	return data, err
}

// GetCategoryVideos gets a list of videos related to a category.
// Name				Type		Required		Description
// page				int			No					The page number to show.
// per_page		int			No					Number of items to show on each page. Max 50.
// query			string	No					Search query.
// sort				string 	No					Technique used to sort the results.
// 						date
// 						alphabetical
// 						videos
// 						followers
// direction	string	No					The direction that the results are sorted.
// 						asc
// 						desc
func (c *Client) GetCategoryVideos(category string, params *Parameters) (*VideoData, error) {
	uri := fmt.Sprintf("/categories/%s/videos", category)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}
	data := &VideoData{}
	err = c.processRequestData(200, resp, data)
	return data, err
}

// CheckIfVideoInCategory checks if a category contains a video
// Name				Type		Required		Description
// page				int			No					The page number to show.
// per_page		int			No					Number of items to show on each page. Max 50.
// query			string	No					Search query.
// sort				string 	No					Technique used to sort the results.
// 						date
// 						alphabetical
// 						videos
// 						followers
// direction	string	No					The direction that the results are sorted.
// 						asc
// 						desc
func (c *Client) CheckIfVideoInCategory(category, video string, params *Parameters) (*VideoDataElement, error) {
	uri := fmt.Sprintf("/categories/%s/videos/%s", category, video)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &VideoDataElement{}
	err = c.processRequestData(200, resp, data)
	return data, err
}
