package gimeo

import (
	"fmt"
	"io/ioutil"

	. "github.com/julianedialkova/gimeo/data"
)

//GetMe gets information about me
func (c *Client) GetMe(params *Parameters) (*User, error) {
	resp, err := c.Get("/me", params)

	if err != nil {
		return nil, err
	}

	data := &User{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

//ChangeMe edits my users
// Name											Type				Required		Description
// videos.privacy.download	boolean			No					Sets the default download setting for all future videos uploaded by this user. If true, the video can be downloaded by any user.
// videos.privacy.add				boolean			No					Sets the default add setting for all future videos uploaded by this user. If true, anyone can add the video to an album, channel, or group.
// videos.privacy.comments	string			No					Sets the default comment setting for all future videos uploaded by this user. It specifies who can comment on the video.
// 													anybody
// 													nobody
// 													contacts
// videos.privacy.view			string 			No					Sets the default view setting for all future videos uploaded by this user. It specifies who can view the video.
// 													anybody
// 													nobody
// 													contacts
// 													password
// 													users
// 													unlisted
// 													disable
// videos.privacy.embed			string 			No					Sets the default embed setting for all future videos uploaded by this user. Whitelist allows you to define all valid embed domains. Check out our docs for adding and removing domains.
//
//													public
//													private
//													whitelistname
// name											string			No					The user's display name
//location									string			No					The user's location
//bio												string			No					The user's bio
func (c *Client) ChangeMe(params *Parameters) (*User, error) {
	resp, err := c.Patch("/me", params)

	if err != nil {
		return nil, err
	}

	data := &User{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

//GetMyAlbums gets a list of user's albums
// Name				Type					Required	Description
// page				int						No				The page number to show.
// per_page		int						No				Number of items to show on each page. Max 50.
// query			string				No				Search query.
// sort				string				No				Technique used to sort the results.
//						date
// 						alphabetical
// 						videos
// 						duration
// direction	string				No				The direction that the results are sorted.
// 						asc
// 						desc
func (c *Client) GetMyAlbums(params *Parameters) (*AlbumData, error) {
	resp, err := c.Get("/me/albums", params)

	if err != nil {
		return nil, err
	}

	data := &AlbumData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// PutInMyAlbums create an Album.
// This method requires a token with the "create" scope.
// Name					Type				Required				Description
// name					string			Yes							The Album title
// description	string			Yes							The Album description
// privacy			string 			No							The Album's privacy level
// 							anybody
// 							password
// password			string			No							Required if privacy=password. The Album's password
// sort					string			No							The default sort order of an Album's videos
//							arranged
//							newest
//							oldest
//							plays
//							comments
//							likes
//							added_first
//							added_last
//							alphabetical
func (c *Client) PutInMyAlbums(params *Parameters) (*AlbumDataElement, error) {
	resp, err := c.Post("/me/albums", params)

	if err != nil {
		return nil, err
	}

	data := &AlbumDataElement{}

	err = c.processRequestData(201, resp, data)
	return data, err
}

// GetMyAlbum gets info on an Album.
func (c *Client) GetMyAlbum(album string, params *Parameters) (*AlbumDataElement, error) {
	uri := fmt.Sprintf("/me/albums/%s", album)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &AlbumDataElement{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// ChangeMyAlbum edits an album
// This method requires a token with the "create" scope.
// Name					Type				Required				Description
// name					string			Yes							The Album title
// description	string			Yes							The Album description
// privacy			string 			No							The Album's privacy level
// 							anybody
// 							password
// sort					string			No							The default sort order of an Album's videos
//							arranged
//							newest
//							oldest
//							plays
//							comments
//							likes
//							added_first
//							added_last
//							alphabetical
func (c *Client) ChangeMyAlbum(album string, params *Parameters) (*AlbumDataElement, error) {
	uri := fmt.Sprintf("/me/albums/%s", album)
	resp, err := c.Patch(uri, params)

	if err != nil {
		return nil, err
	}

	data := &AlbumDataElement{}

	//It actually returns 200
	err = c.processRequestData(200, resp, data)
	return data, err
}

//DeleteMyAlbum deletes an Album.
func (c *Client) DeleteMyAlbum(album string, params *Parameters) error {
	uri := fmt.Sprintf("/me/albums/%s", album)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

//GetMyAlbumVideos gets a list of videos in an album
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
// 										manual
// 										date
// 										alphabetical
// 										plays
// 										likes
// 										comments
// 										duration
// 										modified_time
// direction					string			No						The direction that the results are sorted.
// 										asc
// 										desc
func (c *Client) GetMyAlbumVideos(album string, params *Parameters) (*VideoData, error) {
	uri := fmt.Sprintf("me/albums/%s/videos", album)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &VideoData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

//CheckIfVideoInMyAlbum cheks if an Album contains a video.
func (c *Client) CheckIfVideoInMyAlbum(album, video string, params *Parameters) (*VideoDataElement, error) {
	uri := fmt.Sprintf("me/albums/%s/videos/%s", album, video)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &VideoDataElement{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// PutVideoInMyAlbum adds a video to an Album.
// This method requires a token with the "edit" scope.
func (c *Client) PutVideoInMyAlbum(album, video string, params *Parameters) error {
	uri := fmt.Sprintf("me/albums/%s/videos/%s", album, video)
	resp, err := c.Put(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

// DeleteVideoFromMyAlbum removes a video from an Album.
// This method requires a token with the "edit" scope.
func (c *Client) DeleteVideoFromMyAlbum(album, video string, params *Parameters) error {
	uri := fmt.Sprintf("me/albums/%s/videos/%s", album, video)
	resp, err := c.Delete(uri)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(body)

	err = c.processRequestNoData(resp)
	return err
}

// GetMyAppearances gets all videos that a user appears in.
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
// 										durection
// direction					string			No						The direction that the results are sorted.
// 										asc
// 										desc
func (c *Client) GetMyAppearances(params *Parameters) (*VideoData, error) {
	resp, err := c.Get("/me/appearances", params)

	if err != nil {
		return nil, err
	}

	data := &VideoData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

//GetMyChannels gets a list of channels the user follows
// Name								Type				Required			Description
// page								int					No						The page number to show.
// per_page						int					No						Number of items to show on each page. Max 50.
// query							string			No						Search query.
// filter							string			No						Filter to apply to the results.
// 										moderated
// sort								string			No						Technique used to sort the results.
// 										date
// 										alphabetical
// 										videos
// 										followers
// direction					string			No						The direction that the results are sorted.
// 										asc
// 										desc
func (c *Client) GetMyChannels(params *Parameters) (*ChannelData, error) {
	resp, err := c.Get("/me/channels", params)

	if err != nil {
		return nil, err
	}

	data := &ChannelData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

//MeFollowingChannel checks if a user follows a Channel.
func (c *Client) MeFollowingChannel(channel string, params *Parameters) (bool, error) {
	uri := fmt.Sprintf("me/channels/%s", channel)
	resp, err := c.Get(uri, params)

	if err != nil {
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

//MeFollowChannel subscribes to a Channel.
// This method requires a token with the "interact" scope.
func (c *Client) MeFollowChannel(channel string, params *Parameters) error {
	uri := fmt.Sprintf("me/channels/%s", channel)
	resp, err := c.Put(uri)

	if err != nil {
		return err
	}

	return c.processRequestNoData(resp)
}

// MeUnfollowChannel unsubscribes from a Channel.
// This method requires a token with the "interact" scope.
func (c *Client) MeUnfollowChannel(channel string, params *Parameters) error {

	uri := fmt.Sprintf("me/channels/%s", channel)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	return c.processRequestNoData(resp)

}

// MeFollowingCategory checks if a user follows a Category.
func (c *Client) MeFollowingCategory(category string, params *Parameters) (bool, error) {
	uri := fmt.Sprintf("me/categories/%s", category)
	resp, err := c.Get(uri, params)

	if err != nil {
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

// MeFollowCategory subscribes to a category
// This method requires a token with the "interact" scope.
func (c *Client) MeFollowCategory(category string, params *Parameters) error {
	uri := fmt.Sprintf("me/categories/%s", category)
	resp, err := c.Put(uri)

	if err != nil {
		return err
	}

	return c.processRequestNoData(resp)
}

// MeUnfollowCategory unsubscribes from a category
// This method requires a token with the "interact" scope.
func (c *Client) MeUnfollowCategory(category string, params *Parameters) error {

	uri := fmt.Sprintf("me/categories/%s", category)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	return c.processRequestNoData(resp)

}

//GetMyGroups gets a list of groups the user has joined
// Name								Type				Required			Description
// page								int					No						The page number to show.
// per_page						int					No						Number of items to show on each page. Max 50.
// query							string			No						Search query.
// filter							string			No						Filter to apply to the results.
// 										moderated
// sort								string			No						Technique used to sort the results.
// 										date
// 										alphabetical
// 										videos
// 										followers
// direction					string			No						The direction that the results are sorted.
// 										asc
// 										desc
func (c *Client) GetMyGroups(params *Parameters) (*GroupData, error) {
	resp, err := c.Get("/me/groups", params)

	if err != nil {
		return nil, err
	}

	data := &GroupData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

//MeFollowingGroup checks if an user has joined a group
func (c *Client) MeFollowingGroup(group string, params *Parameters) (bool, error) {
	uri := fmt.Sprintf("me/groups/%s", group)
	resp, err := c.Get(uri, params)

	if err != nil {
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

//MeFollowGroup joins a group
//This method requires a token with the "interact" scope.
func (c *Client) MeFollowGroup(group string, params *Parameters) error {
	uri := fmt.Sprintf("me/groups/%s", group)
	resp, err := c.Put(uri)

	if err != nil {
		return err
	}

	return c.processRequestNoData(resp)
}

//MeUnfollowGroup leaves a group
//This method requires a token with the "interact" scope.
func (c *Client) MeUnfollowGroup(group string, params *Parameters) error {

	uri := fmt.Sprintf("me/groups/%s", group)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	return c.processRequestNoData(resp)

}

//GetMyFeed gets a list of the videos in user's feed.
// Name					Type			Required	Description
// page					int				No				The page number to show.
// per_page			int				No				Number of items to show on each page. Max 50.
// offset				string		Yes				This is necessary for proper pagination. Do not provide this value yourself, just use the pagination links provided in the feed response
// This method requires a token with the "private" scope.
func (c *Client) GetMyFeed(params *Parameters) (*FeedData, error) {
	resp, err := c.Get("/me/feed", params)

	if err != nil {
		return nil, err
	}

	data := &FeedData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

//GetMyFollowers gets a list of user's followers
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
func (c *Client) GetMyFollowers(params *Parameters) (*FollowersData, error) {
	resp, err := c.Get("/me/followers", params)

	if err != nil {
		return nil, err
	}

	data := &FollowersData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

//GetMyFollowing gets a list of the users that a user is following.
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
// 						online
//
func (c *Client) GetMyFollowing(params *Parameters) (*FollowersData, error) {
	resp, err := c.Get("/me/following", params)

	if err != nil {
		return nil, err
	}

	data := &FollowersData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// MeFollowingUser checks if a user follows another user.
func (c *Client) MeFollowingUser(user string, params *Parameters) (bool, error) {
	uri := fmt.Sprintf("me/following/%s", user)
	resp, err := c.Get(uri, params)

	if err != nil {
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

//MeFollowUser follows an user
// This method requires a token with the "interact" scope.
func (c *Client) MeFollowUser(user string, params *Parameters) error {
	uri := fmt.Sprintf("me/following/%s", user)
	resp, err := c.Put(uri)

	if err != nil {
		return err
	}

	return c.processRequestNoData(resp)
}

// MeUnfollowUser unfollows an user
// This method requires a token with the "interact" scope.
func (c *Client) MeUnfollowUser(user string, params *Parameters) error {

	uri := fmt.Sprintf("me/following/%s", user)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	return c.processRequestNoData(resp)

}

//GetMyLikes gets a list of videos that a user likes.
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
func (c *Client) GetMyLikes(params *Parameters) (*LikesData, error) {
	resp, err := c.Get("/me/likes", params)

	if err != nil {
		return nil, err
	}

	data := &LikesData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

//MeLikingVideo checks if a user likes a video.
func (c *Client) MeLikingVideo(video string, params *Parameters) (bool, error) {
	uri := fmt.Sprintf("me/likes/%s", video)
	resp, err := c.Get(uri, params)

	if err != nil {
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

// MeLikeVideo likes a video
// This method requires a token with the "interact" scope.
func (c *Client) MeLikeVideo(video string, params *Parameters) error {
	uri := fmt.Sprintf("me/likes/%s", video)
	resp, err := c.Put(uri)

	if err != nil {
		return err
	}

	return c.processRequestNoData(resp)
}

// MeUnlikeVideo unlikes a video
// This method requires a token with the "interact" scope.
func (c *Client) MeUnlikeVideo(video string, params *Parameters) error {

	uri := fmt.Sprintf("me/likes/%s", video)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	return c.processRequestNoData(resp)

}

// GetMyPictures gets a list of this user's portrait images.
func (c *Client) GetMyPictures(params *Parameters) (*PicturesData, error) {
	resp, err := c.Get("/me/pictures", params)

	if err != nil {
		return nil, err
	}

	data := &PicturesData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// UploadPicture creaetes a new picture resource.
// This method requires a token with the "upload" scope.
func (c *Client) UploadPicture(params *Parameters) (*Picture, error) {
	resp, err := c.Post("/me/pictures", params)

	if err != nil {
		return nil, err
	}

	data := &Picture{}
	err = c.processRequestData(201, resp, data)
	return data, err
}

// MeHasPicture checks if a user has a portrait.
func (c *Client) MeHasPicture(picture string, params *Parameters) (*Picture, error) {
	uri := fmt.Sprintf("me/pictures/%s", picture)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &Picture{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// MeEditPicture edits a portrait
func (c *Client) MeEditPicture(picture string, params *Parameters) (*Picture, error) {
	uri := fmt.Sprintf("me/pictures/%s", picture)
	resp, err := c.Patch(uri, params)

	if err != nil {
		return nil, err
	}

	data := &Picture{}

	err = c.processRequestData(204, resp, data)
	return data, err
}

// DeleteMyPicture removes a portrait from your portrait list.
// This method requires a token with the "delete" scope.
func (c *Client) DeleteMyPicture(picture string, params *Parameters) error {
	uri := fmt.Sprintf("/me/pictures/%s", picture)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

//GetMyPorfolios gets a list of Portfolios created by a user.
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
func (c *Client) GetMyPorfolios(params *Parameters) (*PortfolioData, error) {
	resp, err := c.Get("/me/portfolios", params)

	if err != nil {
		return nil, err
	}

	data := &PortfolioData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// GetMyPortfolio gets a portfolio
func (c *Client) GetMyPortfolio(portfolio string, params *Parameters) (*Portfolio, error) {
	uri := fmt.Sprintf("/me/portfolios/%s", portfolio)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &Portfolio{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

//GetMyPortfolioVideos gets the videos in this Portfolio.
// Name								Type				Required			Description
// page								int					No						The page number to show.
// per_page						int					No						Number of items to show on each page. Max 50.
// sort								string			No						Technique used to sort the results.
// 										date
// 										alphabetical
// 										plays
// 										likes
// 										comments
// 										manual
//										default
// direction					string			No						The direction that the results are sorted.
// 										asc
// 										desc
func (c *Client) GetMyPortfolioVideos(portfolio string, params *Parameters) (*VideoData, error) {
	uri := fmt.Sprintf("/me/portfolios/%s/videos", portfolio)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &VideoData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// MyPortfoliosContaintsVideo checks if a Portfolio contains a video.
func (c *Client) MyPortfoliosContaintsVideo(portfolio, video string, params *Parameters) (bool, error) {
	uri := fmt.Sprintf("me/portfolios/%s/videos/%s", portfolio, video)
	resp, err := c.Get(uri, params)

	if err != nil {
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

// MyPortfoliosAddVideo adds a video to the Portfolio.
// This method requires a token with the "edit" scope.
func (c *Client) MyPortfoliosAddVideo(portfolio, video string, params *Parameters) error {
	uri := fmt.Sprintf("me/portfolios/%s/videos/%s", portfolio, video)
	resp, err := c.Put(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

// MyPortfoliosDeleteVideos removes a video from the Portfolio.
// This method requires a token with the "edit" scope.
func (c *Client) MyPortfoliosDeleteVideos(portfolio, video string, params *Parameters) error {
	uri := fmt.Sprintf("me/portfolios/%s/videos/%s", portfolio, video)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

// GetMyWatchedVideos gets a list of watched videos
func (c *Client) GetMyWatchedVideos(params *Parameters) (*VideoData, error) {
	resp, err := c.Get("/me/watched/videos", params)

	if err != nil {
		return nil, err
	}

	data := &VideoData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// DeleteMyWatchedVideos deletes all watched videos
func (c *Client) DeleteMyWatchedVideos(params *Parameters) error {
	resp, err := c.Delete("/me/watched/videos")

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err

}

// DeleteMyWatchedVideo deletes a watched video
func (c *Client) DeleteMyWatchedVideo(video string, params *Parameters) error {
	uri := fmt.Sprintf("me/watched/videos/%s", video)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

// MeUploadVideo begins the video upload process.
// This method requires a token with the "upload" scope.
func (c *Client) MeUploadVideo(params *Parameters) (*Upload, error) {
	resp, err := c.Post("/me/videos", params)

	if err != nil {
		return nil, err
	}

	data := &Upload{}
	err = c.processRequestData(201, resp, data)
	return data, err
}

// GetMyVideos gets a list of videos uploaded by a user.
// Name								Type					Required	Description
// page								int						No				The page number to show.
// per_page						int						No				Number of items to show on each page. Max 50.
// query							string				No				Search query.
// filter							string				No				Filter to apply to the results.
// 										embeddable
// 										playable
// filter_embeddable	string				No				Required if filter=embeddable. Choose between only videos that are embeddable, and only videos that are not embeddable.
// 										true
// 										false
//
// filter_playable		string				No				Default true. Choose between only videos that are playable, and only videos that are not playable.
// 										true
// 										false
// sort								string  			No				Technique used to sort the results.
// 										date
// 										alphabetical
// 										plays
// 										likes
// 										comments
// 										duration
// 										default
// 										modified_time
// direction					string				No				The direction that the results are sorted.
// 										asc
// 										desc
func (c *Client) GetMyVideos(params *Parameters) (*VideoData, error) {
	uri := fmt.Sprintf("me/videos")
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &VideoData{}

	err = c.processRequestData(200, resp, data)
	if err != nil {
		return nil, err
	}
	return data, err
}

//GetMyVideo checks if a user owns a clip.
func (c *Client) GetMyVideo(video string, params *Parameters) (*VideoDataElement, error) {
	uri := fmt.Sprintf("me/videos/%s", video)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &VideoDataElement{}

	err = c.processRequestData(200, resp, data)
	if err != nil {
		return nil, err
	}
	return data, err
}

//GetWatchLater gets the authenticated user's Watch Later queue.
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
func (c *Client) GetWatchLater(params *Parameters) (*VideoDataElement, error) {
	resp, err := c.Get("/me/watchlater", params)

	if err != nil {
		return nil, err
	}

	data := &VideoDataElement{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

//CheckVideoInMyWatchedlater checks if a video is in the authenticated user's Watch Later queue.
func (c *Client) CheckVideoInMyWatchedlater(video string, params *Parameters) (bool, error) {
	uri := fmt.Sprintf("me/watchlater/%s", video)
	resp, err := c.Get(uri, params)

	if err != nil {
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

// AddWatchLater adds a video to the authenticated user's watch later list.
// This method requires a token with the "interact" scope.
func (c *Client) AddWatchLater(video string, params *Parameters) error {
	uri := fmt.Sprintf("me/watchlater/%s", video)
	resp, err := c.Put(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

// DeleteMyWatchlaterVideo removes a video from your watch later list.
// This method requires a token with the "interact" scope.
func (c *Client) DeleteMyWatchlaterVideo(video string, params *Parameters) error {
	uri := fmt.Sprintf("me/watchlater/%s", video)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}
