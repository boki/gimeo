package gimeo

import (
	"fmt"

	. "github.com/julianedialkova/gimeo/data"
)

func (c *Client) GetVideos(params *Parameters) (*VideoData, error) {
	resp, err := c.Get("/videos", params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &VideoData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) GetVideo(video string, params *Parameters) (*VideoDataElement, error) {
	uri := fmt.Sprintf("/videos/%s", video)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &VideoDataElement{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) ChangeVideo(video string, params *Parameters) (*VideoDataElement, error) {
	uri := fmt.Sprintf("/videos/%s", video)
	resp, err := c.Patch(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &VideoDataElement{}

	err = c.processRequestData(204, resp, data)
	return data, err
}

func (c *Client) DeleteVideo(video string, params *Parameters) error {
	uri := fmt.Sprintf("/videos/%s", video)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

func (c *Client) ReplaceVideo(video string, params *Parameters) error {
	uri := fmt.Sprintf("/videos/%s/files", video)
	resp, err := c.Put(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

func (c *Client) GetVideoCredits(video string, params *Parameters) (*CreditsData, error) {
	uri := fmt.Sprintf("/videos/%s/credits", video)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &CreditsData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) AddCreditToVideo(video string, params *Parameters) (*CreditsDataElement, error) {
	uri := fmt.Sprintf("/videos/%s/credits", video)
	resp, err := c.Post(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &CreditsDataElement{}

	err = c.processRequestData(201, resp, data)
	return data, err
}

func (c *Client) GetRelatedVideos(video string, params *Parameters) (*VideoData, error) {
	uri := fmt.Sprintf("/videos/%s/videos", video)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &VideoData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) GetVideoCategories(video string, params *Parameters) (*CategoryData, error) {
	uri := fmt.Sprintf("/videos/%s/categories", video)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &CategoryData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) AddVideoCategories(video string, params *Parameters) error {
	uri := fmt.Sprintf("/videos/%s/categories", video)
	resp, err := c.Put(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	return c.processRequestNoData(resp)
}

func (c *Client) GetVideoCredit(video, credit string, params *Parameters) (*CreditsDataElement, error) {
	uri := fmt.Sprintf("/videos/%s/credits/%s", video, credit)
	resp, err := c.Get(uri, params)
	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &CreditsDataElement{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) ChangeVideoCredit(video, credit string, params *Parameters) (*Picture, error) {
	uri := fmt.Sprintf("/videos/%s/credits/%s", video, credit)
	resp, err := c.Patch(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &Picture{}

	err = c.processRequestData(204, resp, data)
	return data, err
}

func (c *Client) DeleteVideoCredit(video, credit string, params *Parameters) error {
	uri := fmt.Sprintf("/videos/%s/credits/%s", video, credit)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

func (c *Client) GetVideoComments(video string, params *Parameters) (*CommentData, error) {
	uri := fmt.Sprintf("/videos/%s/comments", video)
	resp, err := c.Get(uri, params)
	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &CommentData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) CommentVideo(video string, params *Parameters) (*CommentDataElement, error) {
	uri := fmt.Sprintf("/videos/%s/comments", video)
	resp, err := c.Post(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &CommentDataElement{}

	err = c.processRequestData(201, resp, data)
	return data, err
}

func (c *Client) ChangeVideoComment(video, comment string, params *Parameters) (*CommentDataElement, error) {
	uri := fmt.Sprintf("/videos/%s/comments/%s", video, comment)
	resp, err := c.Patch(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &CommentDataElement{}

	err = c.processRequestData(204, resp, data)
	return data, err
}

func (c *Client) DeleteVideoComment(video, comment string, params *Parameters) error {
	uri := fmt.Sprintf("/videos/%s/comments/%s", video, comment)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

func (c *Client) GetVideoCommentReplies(video, comment string, params *Parameters) (*CommentData, error) {
	uri := fmt.Sprintf("/videos/%s/comments/%s/replies", video, comment)
	resp, err := c.Get(uri, params)
	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &CommentData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) ReplyToCommentVideo(video, comment string, params *Parameters) (*CommentDataElement, error) {
	uri := fmt.Sprintf("/videos/%s/comments/%s/replies", video, comment)
	resp, err := c.Post(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &CommentDataElement{}

	err = c.processRequestData(201, resp, data)
	return data, err
}

func (c *Client) GetVideoPictures(video string, params *Parameters) (*PicturesData, error) {
	uri := fmt.Sprintf("/videos/%s/pictures", video)
	resp, err := c.Get(uri, params)
	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &PicturesData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) AddPictureToVideo(video string, params *Parameters) (*VideoDataElement, error) {
	uri := fmt.Sprintf("/videos/%s/pictures", video)
	resp, err := c.Post(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &VideoDataElement{}

	err = c.processRequestData(201, resp, data)
	return data, err
}

func (c *Client) GetVideoPicture(video, picture string, params *Parameters) (*Picture, error) {
	uri := fmt.Sprintf("/videos/%s/pictures/%s", video, picture)
	resp, err := c.Get(uri, params)
	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &Picture{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) ChangeVideoPicture(video, picture string, params *Parameters) (*Picture, error) {
	uri := fmt.Sprintf("/videos/%s/pictures/%s", video, picture)
	resp, err := c.Patch(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &Picture{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) DeleteVideoPicture(video string, params *Parameters) error {
	uri := fmt.Sprintf("/videos/%s/pictures/%s", video)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

func (c *Client) GetVideoLikes(video string, params *Parameters) (*LikesData, error) {
	uri := fmt.Sprintf("/videos/%s/likes", video)
	resp, err := c.Get(uri, params)
	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &LikesData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) GetVideoTags(video string, params *Parameters) (*TagsData, error) {
	uri := fmt.Sprintf("/videos/%s/tags", video)
	resp, err := c.Get(uri, params)
	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &TagsData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) CheckVideoTag(video, tag string, params *Parameters) (*TagsData, error) {
	uri := fmt.Sprintf("/videos/%s/tags/%s", video, tag)
	resp, err := c.Get(uri, params)
	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &TagsData{}

	err = c.processRequestData(204, resp, data)
	return data, err
}

func (c *Client) VideoPutTag(video, tag string, params *Parameters) error {
	uri := fmt.Sprintf("/videos/%s/tags/%s", video, tag)
	resp, err := c.Put(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

func (c *Client) DeleteVideoTag(video, tag string, params *Parameters) error {
	uri := fmt.Sprintf("/videos/%s/tags/%s", video, tag)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

func (c *Client) GetVideoUsers(video string, params *Parameters) (*UserData, error) {
	uri := fmt.Sprintf("/videos/%s/privacy/users", video)
	resp, err := c.Get(uri, params)
	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &UserData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) AddvideoUsers(video string, params *Parameters) (*UserData, error) {
	uri := fmt.Sprintf("/videos/%s/privacy/users", video)
	resp, err := c.Put(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}
	data := &UserData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) AddVideoUser(video, user string, params *Parameters) error {
	uri := fmt.Sprintf("/videos/%s/privacy/users/%s", video, user)
	resp, err := c.Put(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

func (c *Client) DeleteVideoUser(video, user string, params *Parameters) error {
	uri := fmt.Sprintf("/videos/%s/privacy/users/%s", video, user)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}
