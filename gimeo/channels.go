package gimeo

import (
	"fmt"

	. "github.com/julianedialkova/gimeo/data"
)

// GetChannels gets a list of all Channels.
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
//filter			string				No				Filter to apply to the results.
// 						feature
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

//CreateChannel creates a new Channel.
// Name					Type			Required			Description
// name					string		Yes						The name of the new Channel
// description	string		Yes						The description of the new Channel
// privacy			string		Yes						The privacy level of the new Channel
// 							anybody
// 							user
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

// GetChannel gets a Channel.
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

// PatchChannel edits a Channel's information
// Name					Type			Required			Description
// name					string		No						The Channel's new name
// description	string		No						The Channel's new description
// privacy			string		No						The Channel's new privacy level
// 							anybody
// 							users
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

//DeleteChannel deletes a Channel
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

// GetChannelUsers gets a list of users who follow a Channel.
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
//filter			string				No				Filter to apply to the results.
// 						featured
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

// GetChannelVideos gets a list of videos in a Channel.
// Name								Type					Required	Description
// page								int						No				The page number to show.
// per_page						int						No				Number of items to show on each page. Max 50.
// query							string				No				Search query.
// sort								string				Но				No	Technique used to sort the results.
//										date
//							  		alphabetical
//						   			plays
//					  				likes
//										comments
//										duration
//										added
//										modified_time
//										manual
// direction					string				Но				No	The direction that the results are sorted.
// 										asc
// 										desc
// filter							string				No				Filter to apply to the results.
// 										embeddable
// filter_embeddable	string				No				Required if filter=embeddable. Choose between only videos that are embeddable, and only videos that are not embeddable.
//										true
//										false
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

// CheckIfVideoInChannel checks if this Channel contains a video.
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

// PutVideoInChannel adds a video to a Channel.
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

// DeleteVideoFromChannel removes a video from a Channel.
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
