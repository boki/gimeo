package gimeo

import (
	"fmt"
	"io/ioutil"

	. "github.com/julianedialkova/gimeo/data"
)

func (c *Client) GetMe(params *Parameters) (*User, error) {
	resp, err := c.Get("/me", params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &User{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) ChangeMe(params *Parameters) (*User, error) {
	resp, err := c.Patch("/me", params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &User{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) GetMyAlbums(params *Parameters) (*AlbumData, error) {
	resp, err := c.Get("/me/albums", params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &AlbumData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) PutInMyAlbums(params *Parameters) (*AlbumDataElement, error) {
	resp, err := c.Post("/me/albums", params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &AlbumDataElement{}

	err = c.processRequestData(201, resp, data)
	return data, err
}

func (c *Client) GetMyAlbum(album string, params *Parameters) (*AlbumDataElement, error) {
	uri := fmt.Sprintf("/me/albums/%s", album)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &AlbumDataElement{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) ChangeMyAlbum(album string, params *Parameters) (*AlbumDataElement, error) {
	uri := fmt.Sprintf("/me/albums/%s", album)
	resp, err := c.Patch(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &AlbumDataElement{}

	//It actually returns 200
	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) DeleteMyAlbum(album string, params *Parameters) error {
	uri := fmt.Sprintf("/me/albums/%s", album)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

func (c *Client) GetMyAlbumVideos(album string, params *Parameters) (*VideoData, error) {
	uri := fmt.Sprintf("me/albums/%s/videos", album)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &VideoData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) CheckIfVideoInMyAlbum(album, video string, params *Parameters) (*VideoDataElement, error) {
	uri := fmt.Sprintf("me/albums/%s/videos/%s", album, video)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &VideoDataElement{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) PutVideoInMyAlbum(album, video string, params *Parameters) error {
	uri := fmt.Sprintf("me/albums/%s/videos/%s", album, video)
	resp, err := c.Put(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

func (c *Client) DeleteVideoFromMyAlbum(album, video string, params *Parameters) error {
	uri := fmt.Sprintf("me/albums/%s/videos/%s", album, video)
	resp, err := c.Delete(uri)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Could not read requests body")
		return err
	}

	fmt.Println(body)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

func (c *Client) GetMyAppearances(params *Parameters) (*VideoData, error) {
	resp, err := c.Get("/me/appearances", params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &VideoData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) GetMyChannels(params *Parameters) (*ChannelData, error) {
	resp, err := c.Get("/me/channels", params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &ChannelData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) MeFollowingChannel(channel string, params *Parameters) (bool, error) {
	uri := fmt.Sprintf("me/channels/%s", channel)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return false, err
	}

	if resp.StatusCode == 204 {
		return true, nil
	} else if resp.StatusCode == 404 {
		return false, nil
	} else {
		fmt.Printf("Request returned %s.\n", resp.Status)
		return false, fmt.Errorf("Request returned %s.\n", resp.Status)
	}
}

func (c *Client) MeFollowChannel(channel string, params *Parameters) error {
	uri := fmt.Sprintf("me/channels/%s", channel)
	resp, err := c.Put(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	return c.processRequestNoData(resp)
}

func (c *Client) MeUnfollowChannel(channel string, params *Parameters) error {

	uri := fmt.Sprintf("me/channels/%s", channel)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	return c.processRequestNoData(resp)

}

func (c *Client) MeFollowingCategory(category string, params *Parameters) (bool, error) {
	uri := fmt.Sprintf("me/categories/%s", category)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return false, err
	}

	if resp.StatusCode == 204 {
		return true, nil
	} else if resp.StatusCode == 404 {
		return false, nil
	} else {
		fmt.Printf("Request returned %s.\n", resp.Status)
		return false, fmt.Errorf("Request returned %s.\n", resp.Status)
	}
}

func (c *Client) MeFollowCategory(category string, params *Parameters) error {
	uri := fmt.Sprintf("me/categories/%s", category)
	resp, err := c.Put(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	return c.processRequestNoData(resp)
}

func (c *Client) MeUnfollowCategory(category string, params *Parameters) error {

	uri := fmt.Sprintf("me/categories/%s", category)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	return c.processRequestNoData(resp)

}

func (c *Client) GetMyGroups(params *Parameters) (*GroupData, error) {
	resp, err := c.Get("/me/groups", params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &GroupData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) MeFollowingGroup(group string, params *Parameters) (bool, error) {
	uri := fmt.Sprintf("me/groups/%s", group)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return false, err
	}

	if resp.StatusCode == 204 {
		return true, nil
	} else if resp.StatusCode == 404 {
		return false, nil
	} else {
		fmt.Printf("Request returned %s.\n", resp.Status)
		return false, fmt.Errorf("Request returned %s.\n", resp.Status)
	}
}

func (c *Client) MeFollowGroup(group string, params *Parameters) error {
	uri := fmt.Sprintf("me/groups/%s", group)
	resp, err := c.Put(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	return c.processRequestNoData(resp)
}

func (c *Client) MeUnfollowGroup(group string, params *Parameters) error {

	uri := fmt.Sprintf("me/groups/%s", group)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	return c.processRequestNoData(resp)

}

func (c *Client) GetMyFeed(params *Parameters) (*FeedData, error) {
	resp, err := c.Get("/me/feed", params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &FeedData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) GetMyFollowers(params *Parameters) (*FollowersData, error) {
	resp, err := c.Get("/me/followers", params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &FollowersData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) GetMyFollowing(params *Parameters) (*FollowersData, error) {
	resp, err := c.Get("/me/following", params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &FollowersData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) MeFollowingUser(user string, params *Parameters) (bool, error) {
	uri := fmt.Sprintf("me/following/%s", user)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return false, err
	}

	if resp.StatusCode == 204 {
		return true, nil
	} else if resp.StatusCode == 404 {
		return false, nil
	} else {
		fmt.Printf("Request returned %s.\n", resp.Status)
		return false, fmt.Errorf("Request returned %s.\n", resp.Status)
	}
}

func (c *Client) MeFollowUser(user string, params *Parameters) error {
	uri := fmt.Sprintf("me/following/%s", user)
	resp, err := c.Put(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	return c.processRequestNoData(resp)
}

func (c *Client) MeUnfollowUser(user string, params *Parameters) error {

	uri := fmt.Sprintf("me/following/%s", user)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	return c.processRequestNoData(resp)

}

func (c *Client) GetMyLikes(params *Parameters) (*LikesData, error) {
	resp, err := c.Get("/me/likes", params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &LikesData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) MeLikingVideo(video string, params *Parameters) (bool, error) {
	uri := fmt.Sprintf("me/likes/%s", video)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return false, err
	}

	if resp.StatusCode == 204 {
		return true, nil
	} else if resp.StatusCode == 404 {
		return false, nil
	} else {
		fmt.Printf("Request returned %s.\n", resp.Status)
		return false, fmt.Errorf("Request returned %s.\n", resp.Status)
	}
}

func (c *Client) MeLikeVideo(video string, params *Parameters) error {
	uri := fmt.Sprintf("me/likes/%s", video)
	resp, err := c.Put(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	return c.processRequestNoData(resp)
}

func (c *Client) MeUnlikeVideo(video string, params *Parameters) error {

	uri := fmt.Sprintf("me/likes/%s", video)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	return c.processRequestNoData(resp)

}

func (c *Client) GetMyPictures(params *Parameters) (*PicturesData, error) {
	resp, err := c.Get("/me/pictures", params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &PicturesData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) UploadPicture(params *Parameters) (*Picture, error) {
	resp, err := c.Post("/me/pictures", params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &Picture{}
	err = c.processRequestData(201, resp, data)
	return data, err
}

func (c *Client) MeHasPicture(picture string, params *Parameters) (*Picture, error) {
	uri := fmt.Sprintf("me/pictures/%s", picture)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &Picture{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) MeEditPicture(picture string, params *Parameters) (*Picture, error) {
	uri := fmt.Sprintf("me/pictures/%s", picture)
	resp, err := c.Patch(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &Picture{}

	err = c.processRequestData(204, resp, data)
	return data, err
}

func (c *Client) DeleteMyPicture(picture string, params *Parameters) error {
	uri := fmt.Sprintf("/me/pictures/%s", picture)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

func (c *Client) GetMyPorfolios(params *Parameters) (*PortfolioData, error) {
	resp, err := c.Get("/me/portfolios", params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &PortfolioData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) GetMyPortfolio(portfolio string, params *Parameters) (*Portfolio, error) {
	uri := fmt.Sprintf("/me/portfolios/%s", portfolio)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &Portfolio{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) GetMyPortfolioVideos(portfolio string, params *Parameters) (*VideoData, error) {
	uri := fmt.Sprintf("/me/portfolios/%s/videos", portfolio)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &VideoData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) MyPortfoliosContaintsVideo(portfolio, video string, params *Parameters) (bool, error) {
	uri := fmt.Sprintf("me/portfolios/%s/videos/%s", portfolio, video)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return false, err
	}

	if resp.StatusCode == 204 {
		return true, nil
	} else if resp.StatusCode == 404 {
		return false, nil
	} else {
		fmt.Printf("Request returned %s.\n", resp.Status)
		return false, fmt.Errorf("Request returned %s.\n", resp.Status)
	}
}

func (c *Client) MyPortfoliosAddVideo(portfolio, video string, params *Parameters) error {
	uri := fmt.Sprintf("me/portfolios/%s/videos/%s", portfolio, video)
	resp, err := c.Put(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

func (c *Client) MyPortfoliosDeleteVideos(portfolio, video string, params *Parameters) error {
	uri := fmt.Sprintf("me/portfolios/%s/videos/%s", portfolio, video)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

func (c *Client) GetMyWatchedVideos(params *Parameters) (*VideoData, error) {
	resp, err := c.Get("/me/watched/videos", params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &VideoData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) DeleteMyWatchedVideos(params *Parameters) error {
	resp, err := c.Delete("/me/watched/videos")

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err

}

func (c *Client) DeleteMyWatchedVideo(video string, params *Parameters) error {
	uri := fmt.Sprintf("me/watched/videos/%s", video)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

func (c *Client) MeUploadVideo(params *Parameters) (*Upload, error) {
	resp, err := c.Post("/me/videos", params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &Upload{}
	err = c.processRequestData(201, resp, data)
	return data, err
}

func (c *Client) GetMyVideo(video string, params *Parameters) (*VideoDataElement, error) {
	uri := fmt.Sprintf("me/videos/%s", video)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &VideoDataElement{}

	err = c.processRequestData(200, resp, data)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (c *Client) GetWatchLater(params *Parameters) (*VideoDataElement, error) {
	resp, err := c.Get("/me/watchlater", params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &VideoDataElement{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) CheckVideoInMyWatchedlater(video string, params *Parameters) (bool, error) {
	uri := fmt.Sprintf("me/watchlater/%s", video)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return false, err
	}

	if resp.StatusCode == 204 {
		return true, nil
	} else if resp.StatusCode == 404 {
		return false, nil
	} else {
		fmt.Printf("Request returned %s.\n", resp.Status)
		return false, fmt.Errorf("Request returned %s.\n", resp.Status)
	}
}

func (c *Client) AddWatchLater(video string, params *Parameters) error {
	uri := fmt.Sprintf("me/watchlater/%s", video)
	resp, err := c.Put(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

func (c *Client) DeleteMyWatchlaterVideo(video string, params *Parameters) error {
	uri := fmt.Sprintf("me/watchlater/%s", video)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}
