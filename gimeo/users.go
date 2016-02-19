package gimeo

import (
	"fmt"

	. "github.com/julianedialkova/gimeo/data"
)

func (c *Client) GetUsers(params *Parameters) (*UserData, error) {
	resp, err := c.Get("/users", params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &UserData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) GetUser(user string, params *Parameters) (*User, error) {
	uri := fmt.Sprintf("/users/%s", user)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &User{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) ChangeUser(user string, params *Parameters) (*User, error) {

	uri := fmt.Sprintf("/users/%s", user)
	resp, err := c.Patch(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &User{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) GetUserAlbums(user string, params *Parameters) (*AlbumData, error) {
	uri := fmt.Sprintf("/users/%s/albums", user)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &AlbumData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) PutInUserAlbums(user string, params *Parameters) (*AlbumDataElement, error) {
	uri := fmt.Sprintf("/users/%s/albums", user)
	resp, err := c.Post(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &AlbumDataElement{}

	err = c.processRequestData(201, resp, data)
	return data, err
}

func (c *Client) GetUserAlbum(user string, album string, params *Parameters) (*AlbumDataElement, error) {
	uri := fmt.Sprintf("/users/%s/albums/%s", user, album)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &AlbumDataElement{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) ChangeUserAlbum(user string, album string, params *Parameters) (*AlbumDataElement, error) {
	uri := fmt.Sprintf("/users/%s/albums/%s", user, album)
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

func (c *Client) DeleteUserAlbum(user, album string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/albums/%s", user, album)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

func (c *Client) GetUserAlbumVideos(user, album string, params *Parameters) (*VideoData, error) {
	uri := fmt.Sprintf("/users/%s/albums/%s/videos", user, album)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &VideoData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) CheckIfVideoInUserAlbum(user, album, video string, params *Parameters) (*VideoDataElement, error) {
	uri := fmt.Sprintf("/users/%s/albums/%s/videos/%s", user, album, video)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &VideoDataElement{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) PutVideoInUserAlbum(user, album, video string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/albums/%s/videos/%s", user, album, video)
	resp, err := c.Put(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

func (c *Client) DeleteVideoFromUserAlbum(user, album, video string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/albums/%s/videos/%s", user, album, video)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not read requests body")
		return err
	}

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

func (c *Client) GetUserAppearances(user string, params *Parameters) (*VideoData, error) {
	uri := fmt.Sprintf("/users/%s/appearances", user)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &VideoData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) GetUserChannels(user string, params *Parameters) (*ChannelData, error) {
	uri := fmt.Sprintf("/users/%s/channels", user)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &ChannelData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) UserFollowingChannel(user, channel string, params *Parameters) (bool, error) {
	uri := fmt.Sprintf("/users/%s/channels/%s", user, channel)
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

func (c *Client) UserFollowChannel(user, channel string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/channels/%s", user, channel)
	resp, err := c.Put(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	return c.processRequestNoData(resp)
}

func (c *Client) UserUnfollowChannel(user, channel string, params *Parameters) error {

	uri := fmt.Sprintf("/users/%s/channels/%s", user, channel)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	return c.processRequestNoData(resp)

}

func (c *Client) UserFollowingCategory(user, category string, params *Parameters) (bool, error) {
	uri := fmt.Sprintf("/users/%s/categories/%s", user, category)
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

func (c *Client) UserFollowCategory(user, category string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/categories/%s", user, category)
	resp, err := c.Put(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	return c.processRequestNoData(resp)
}

func (c *Client) UserUnfollowCategory(user, category string, params *Parameters) error {

	uri := fmt.Sprintf("/users/%s/categories/%s", user, category)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	return c.processRequestNoData(resp)

}

func (c *Client) UserMyGroups(user string, params *Parameters) (*GroupData, error) {
	uri := fmt.Sprintf("/users/%s/groups", user)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &GroupData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) UserFollowingGroup(user, group string, params *Parameters) (bool, error) {
	uri := fmt.Sprintf("/users/%s/groups/%s", user, group)
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

func (c *Client) UserFollowGroup(user, group string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/groups/%s", user, group)
	resp, err := c.Put(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	return c.processRequestNoData(resp)
}

func (c *Client) UserUnfollowGroup(user, group string, params *Parameters) error {

	uri := fmt.Sprintf("/users/%s/groups/%s", user, group)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	return c.processRequestNoData(resp)

}

func (c *Client) GetUserFeed(user string, params *Parameters) (*FeedData, error) {
	uri := fmt.Sprintf("/users/%s/feed", user)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &FeedData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) GetUserFollowers(user string, params *Parameters) (*FollowersData, error) {
	uri := fmt.Sprintf("/users/%s/followers", user)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &FollowersData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) GetUserFollowing(user string, params *Parameters) (*FollowersData, error) {
	uri := fmt.Sprintf("/users/%s/following", user)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &FollowersData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) UserFollowingUser(user, followed string, params *Parameters) (bool, error) {
	uri := fmt.Sprintf("/users/%s/following/%s", user, followed)
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

func (c *Client) UserFollowUser(user, followed string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/following/%s", user, followed)
	resp, err := c.Put(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	return c.processRequestNoData(resp)
}

func (c *Client) UserUnfollowUser(user, followed string, params *Parameters) error {

	uri := fmt.Sprintf("/users/%s/following/%s", user, followed)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	return c.processRequestNoData(resp)

}

func (c *Client) GetUserLikes(user string, params *Parameters) (*LikesData, error) {
	uri := fmt.Sprintf("/users/%s/likes", user)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &LikesData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) UserLikingVideo(user, video string, params *Parameters) (bool, error) {
	uri := fmt.Sprintf("/users/%s/likes/%s", user, video)
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

func (c *Client) UserLikeVideo(user, video string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/likes/%s", user, video)
	resp, err := c.Put(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	return c.processRequestNoData(resp)
}

func (c *Client) UserUnlikeVideo(user, video string, params *Parameters) error {

	uri := fmt.Sprintf("/users/%s/likes/%s", user, video)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	return c.processRequestNoData(resp)

}

func (c *Client) GetUserPictures(user string, params *Parameters) (*PicturesData, error) {
	uri := fmt.Sprintf("/users/%s/pictures", user)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &PicturesData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) UserUploadPicture(user string, params *Parameters) (*Picture, error) {
	uri := fmt.Sprintf("/users/%s/pictures", user)
	resp, err := c.Post(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &Picture{}

	err = c.processRequestData(201, resp, data)
	return data, err
}

func (c *Client) UserHasPicture(user, picture string, params *Parameters) (*Picture, error) {
	uri := fmt.Sprintf("/users/%s/pictures/%s", user, picture)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &Picture{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) UserEditPicture(user, picture string, params *Parameters) (*Picture, error) {
	uri := fmt.Sprintf("/users/%s/pictures/%s", user, picture)
	resp, err := c.Patch(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &Picture{}

	err = c.processRequestData(204, resp, data)
	return data, err
}

func (c *Client) DeleteUserPicture(user, picture string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/pictures/%s", user, picture)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

func (c *Client) GetUserPorfolios(user string, params *Parameters) (*PortfolioData, error) {
	uri := fmt.Sprintf("/users/%s/portfolios", user)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &PortfolioData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) GetUserPortfolio(user, portfolio string, params *Parameters) (*Portfolio, error) {
	uri := fmt.Sprintf("/users/%s/portfolios/%s", user, portfolio)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &Portfolio{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) GetUserPortfolioVideos(user, portfolio string, params *Parameters) (*VideoData, error) {
	uri := fmt.Sprintf("/users/%s/portfolios/%s/videos", user, portfolio)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &VideoData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) UserPortfoliosContaintsVideo(user, portfolio, video string, params *Parameters) (bool, error) {
	uri := fmt.Sprintf("/users/%s/portfolios/%s/videos/%s", user, portfolio, video)
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

func (c *Client) UserPortfoliosAddVideo(user, portfolio, video string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/portfolios/%s/videos/%s", user, portfolio, video)
	resp, err := c.Put(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

func (c *Client) UserPortfoliosDeleteVideos(user, portfolio, video string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/portfolios/%s/videos/%s", user, portfolio, video)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

func (c *Client) GetUserWatchedVideos(user string, params *Parameters) (*VideoData, error) {
	uri := fmt.Sprintf("/users/%s/watched/videos", user)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &VideoData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) DeleteUserWatchedVideos(user string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/watched/videos", user)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err

}

func (c *Client) DeleteUserWatchedVideo(user, video string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/watched/videos/%s", user, video)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

func (c *Client) UserUploadVideo(user string, params *Parameters) (*AlbumDataElement, error) {
	uri := fmt.Sprintf("/users/%s/videos", user)
	resp, err := c.Post(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &AlbumDataElement{}

	err = c.processRequestData(201, resp, data)
	return data, err
}

func (c *Client) GetUserVideo(user, video string, params *Parameters) (*VideoDataElement, error) {
	uri := fmt.Sprintf("/users/%s/videos/%s", user, video)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &VideoDataElement{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) GetUserWatchLater(user string, params *Parameters) (*VideoDataElement, error) {
	uri := fmt.Sprintf("/users/%s/watchlater", user)
	resp, err := c.Get(uri, params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &VideoDataElement{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

func (c *Client) CheckVideoInUserWatchedlater(user, video string, params *Parameters) (bool, error) {
	uri := fmt.Sprintf("/users/%s/watchlater/%s", user, video)
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

func (c *Client) UserAddWatchLater(user, video string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/watchlater/%s", user, video)
	resp, err := c.Put(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

func (c *Client) DeleteUserWatchlaterVideo(user, video string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/watchlater/%s", user, video)
	resp, err := c.Delete(uri)

	if err != nil {
		fmt.Println("Could not execute request")
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}
