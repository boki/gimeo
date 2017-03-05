package gimeo

import (
	"fmt"

	. "github.com/julianedialkova/gimeo/data"
)

//GetUsers searches for users
// Name				Type		Required	Description
// page				int			No					The page number to show.
// per_page		int			No					Number of items to show on each page. Max 50.
// query			string	Yes					Search query.
// sort				string	No					Technique used to sort the results.
// 						relevant
// 						date
// 						alphabetical
// direction	string 	No					The direction that the results are sorted.
// 						asc
// 						desc
func (c *Client) GetUsers(params *Parameters) (*UserData, error) {
	resp, err := c.Get("/users", params)

	if err != nil {
		return nil, err
	}

	data := &UserData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

//GetUser gets information for a user
func (c *Client) GetUser(user string, params *Parameters) (*User, error) {
	uri := fmt.Sprintf("/users/%s", user)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &User{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

//ChangeUser edits an user
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
func (c *Client) ChangeUser(user string, params *Parameters) (*User, error) {

	uri := fmt.Sprintf("/users/%s", user)
	resp, err := c.Patch(uri, params)

	if err != nil {
		return nil, err
	}

	data := &User{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

//GetUserAlbums gets a list of user's albums
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
func (c *Client) GetUserAlbums(user string, params *Parameters) (*AlbumData, error) {
	uri := fmt.Sprintf("/users/%s/albums", user)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &AlbumData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// PutInUserAlbums create an Album.
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
func (c *Client) PutInUserAlbums(user string, params *Parameters) (*AlbumDataElement, error) {
	uri := fmt.Sprintf("/users/%s/albums", user)
	resp, err := c.Post(uri, params)

	if err != nil {
		return nil, err
	}

	data := &AlbumDataElement{}

	err = c.processRequestData(201, resp, data)
	return data, err
}

// GetUserAlbum gets info on an Album.
func (c *Client) GetUserAlbum(user string, album string, params *Parameters) (*AlbumDataElement, error) {
	uri := fmt.Sprintf("/users/%s/albums/%s", user, album)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &AlbumDataElement{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// ChangeUserAlbum edits an album
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
func (c *Client) ChangeUserAlbum(user string, album string, params *Parameters) (*AlbumDataElement, error) {
	uri := fmt.Sprintf("/users/%s/albums/%s", user, album)
	resp, err := c.Patch(uri, params)

	if err != nil {
		return nil, err
	}

	data := &AlbumDataElement{}

	//It actually returns 200
	err = c.processRequestData(200, resp, data)
	return data, err
}

//DeleteUserAlbum deletes an Album.
func (c *Client) DeleteUserAlbum(user, album string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/albums/%s", user, album)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

//GetUserAlbumVideos gets a list of videos in an album
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
func (c *Client) GetUserAlbumVideos(user, album string, params *Parameters) (*VideoData, error) {
	uri := fmt.Sprintf("/users/%s/albums/%s/videos", user, album)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &VideoData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

//CheckIfVideoInUserAlbum cheks if an Album contains a video.
func (c *Client) CheckIfVideoInUserAlbum(user, album, video string, params *Parameters) (*VideoDataElement, error) {
	uri := fmt.Sprintf("/users/%s/albums/%s/videos/%s", user, album, video)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &VideoDataElement{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// PutVideoInUserAlbum adds a video to an Album.
// This method requires a token with the "edit" scope.
func (c *Client) PutVideoInUserAlbum(user, album, video string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/albums/%s/videos/%s", user, album, video)
	resp, err := c.Put(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

// DeleteVideoFromUserAlbum removes a video from an Album.
// This method requires a token with the "edit" scope.
func (c *Client) DeleteVideoFromUserAlbum(user, album, video string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/albums/%s/videos/%s", user, album, video)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

// GetUserAppearances gets all videos that a user appears in.
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
func (c *Client) GetUserAppearances(user string, params *Parameters) (*VideoData, error) {
	uri := fmt.Sprintf("/users/%s/appearances", user)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &VideoData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

//GetUserChannels gets a list of channels the user follows
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
func (c *Client) GetUserChannels(user string, params *Parameters) (*ChannelData, error) {
	uri := fmt.Sprintf("/users/%s/channels", user)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &ChannelData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

//UserFollowingChannel checks if a user follows a Channel.
func (c *Client) UserFollowingChannel(user, channel string, params *Parameters) (bool, error) {
	uri := fmt.Sprintf("/users/%s/channels/%s", user, channel)
	resp, err := c.Get(uri, params)

	if err != nil {
		return false, err
	}

	if resp.StatusCode == 204 {
		return true, nil
	} else if resp.StatusCode == 404 {
		return false, nil
	}
	c.log.Printf("Request returned %s.\n", resp.Status)
	return false, fmt.Errorf("request returned %s", resp.Status)
}

//UserFollowChannel subscribes to a Channel.
//This method requires a token with the "interact" scope.
func (c *Client) UserFollowChannel(user, channel string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/channels/%s", user, channel)
	resp, err := c.Put(uri)

	if err != nil {
		return err
	}

	return c.processRequestNoData(resp)
}

//UserUnfollowChannel unsubscribes from a Channel.
//This method requires a token with the "interact" scope.
func (c *Client) UserUnfollowChannel(user, channel string, params *Parameters) error {

	uri := fmt.Sprintf("/users/%s/channels/%s", user, channel)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	return c.processRequestNoData(resp)

}

// UserFollowingCategory checks if a user follows a Category.
func (c *Client) UserFollowingCategory(user, category string, params *Parameters) (bool, error) {
	uri := fmt.Sprintf("/users/%s/categories/%s", user, category)
	resp, err := c.Get(uri, params)

	if err != nil {
		return false, err
	}

	if resp.StatusCode == 204 {
		return true, nil
	} else if resp.StatusCode == 404 {
		return false, nil
	}
	c.log.Printf("Request returned %s.\n", resp.Status)
	return false, fmt.Errorf("request returned %s", resp.Status)
}

//UserFollowCategory subscribes to a category
//This method requires a token with the "interact" scope.
func (c *Client) UserFollowCategory(user, category string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/categories/%s", user, category)
	resp, err := c.Put(uri)

	if err != nil {
		return err
	}

	return c.processRequestNoData(resp)
}

//UserUnfollowCategory unsubscribes from a category
//This method requires a token with the "interact" scope.
func (c *Client) UserUnfollowCategory(user, category string, params *Parameters) error {

	uri := fmt.Sprintf("/users/%s/categories/%s", user, category)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	return c.processRequestNoData(resp)

}

//UserMyGroups gets a list of groups the user has joined
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
func (c *Client) UserMyGroups(user string, params *Parameters) (*GroupData, error) {
	uri := fmt.Sprintf("/users/%s/groups", user)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &GroupData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

//UserFollowingGroup checks if an user has joined a group
func (c *Client) UserFollowingGroup(user, group string, params *Parameters) (bool, error) {
	uri := fmt.Sprintf("/users/%s/groups/%s", user, group)
	resp, err := c.Get(uri, params)

	if err != nil {
		return false, err
	}

	if resp.StatusCode == 204 {
		return true, nil
	} else if resp.StatusCode == 404 {
		return false, nil
	}
	c.log.Printf("Request returned %s.\n", resp.Status)
	return false, fmt.Errorf("request returned %s", resp.Status)
}

//UserFollowGroup joins a group
//This method requires a token with the "interact" scope.
func (c *Client) UserFollowGroup(user, group string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/groups/%s", user, group)
	resp, err := c.Put(uri)

	if err != nil {
		return err
	}

	return c.processRequestNoData(resp)
}

//UserUnfollowGroup leaves a group
//This method requires a token with the "interact" scope.
func (c *Client) UserUnfollowGroup(user, group string, params *Parameters) error {

	uri := fmt.Sprintf("/users/%s/groups/%s", user, group)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	return c.processRequestNoData(resp)

}

//GetUserFeed gets a list of the videos in user's feed.
// Name					Type			Required	Description
// page					int				No				The page number to show.
// per_page			int				No				Number of items to show on each page. Max 50.
// offset				string		Yes				This is necessary for proper pagination. Do not provide this value yourself, just use the pagination links provided in the feed response
// This method requires a token with the "private" scope.
func (c *Client) GetUserFeed(user string, params *Parameters) (*FeedData, error) {
	uri := fmt.Sprintf("/users/%s/feed", user)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &FeedData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

//GetUserFollowers gets a list of user's followers
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
func (c *Client) GetUserFollowers(user string, params *Parameters) (*FollowersData, error) {
	uri := fmt.Sprintf("/users/%s/followers", user)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &FollowersData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

//GetUserFollowing gets a list of the users that a user is following.
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
func (c *Client) GetUserFollowing(user string, params *Parameters) (*FollowersData, error) {
	uri := fmt.Sprintf("/users/%s/following", user)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &FollowersData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// UserFollowingUser checks if a user follows another user.
func (c *Client) UserFollowingUser(user, followed string, params *Parameters) (bool, error) {
	uri := fmt.Sprintf("/users/%s/following/%s", user, followed)
	resp, err := c.Get(uri, params)

	if err != nil {
		return false, err
	}

	if resp.StatusCode == 204 {
		return true, nil
	} else if resp.StatusCode == 404 {
		return false, nil
	} else {
	}
	c.log.Printf("Request returned %s.\n", resp.Status)
	return false, fmt.Errorf("request returned %s", resp.Status)
}

//UserFollowUser follows an user
// his method requires a token with the "interact" scope.
func (c *Client) UserFollowUser(user, followed string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/following/%s", user, followed)
	resp, err := c.Put(uri)

	if err != nil {
		return err
	}

	return c.processRequestNoData(resp)
}

//UserUnfollowUser unfollows an user
// This method requires a token with the "interact" scope.
func (c *Client) UserUnfollowUser(user, followed string, params *Parameters) error {

	uri := fmt.Sprintf("/users/%s/following/%s", user, followed)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	return c.processRequestNoData(resp)

}

//GetUserLikes gets a list of videos that a user likes.
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
func (c *Client) GetUserLikes(user string, params *Parameters) (*LikesData, error) {
	uri := fmt.Sprintf("/users/%s/likes", user)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &LikesData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

//UserLikingVideo checks if a user likes a video.
func (c *Client) UserLikingVideo(user, video string, params *Parameters) (bool, error) {
	uri := fmt.Sprintf("/users/%s/likes/%s", user, video)
	resp, err := c.Get(uri, params)

	if err != nil {
		return false, err
	}

	if resp.StatusCode == 204 {
		return true, nil
	} else if resp.StatusCode == 404 {
		return false, nil
	}
	c.log.Printf("Request returned %s.\n", resp.Status)
	return false, fmt.Errorf("request returned %s", resp.Status)
}

// UserLikeVideo likes a video
// This method requires a token with the "interact" scope.
func (c *Client) UserLikeVideo(user, video string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/likes/%s", user, video)
	resp, err := c.Put(uri)

	if err != nil {
		return err
	}

	return c.processRequestNoData(resp)
}

// UserUnlikeVideo unlikes a video
// This method requires a token with the "interact" scope.
func (c *Client) UserUnlikeVideo(user, video string, params *Parameters) error {

	uri := fmt.Sprintf("/users/%s/likes/%s", user, video)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	return c.processRequestNoData(resp)

}

// GetUserPictures gets a list of this user's portrait images.
func (c *Client) GetUserPictures(user string, params *Parameters) (*PicturesData, error) {
	uri := fmt.Sprintf("/users/%s/pictures", user)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &PicturesData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// UserUploadPicture creaetes a new picture resource.
// This method requires a token with the "upload" scope.
func (c *Client) UserUploadPicture(user string, params *Parameters) (*Picture, error) {
	uri := fmt.Sprintf("/users/%s/pictures", user)
	resp, err := c.Post(uri, params)

	if err != nil {
		return nil, err
	}

	data := &Picture{}

	err = c.processRequestData(201, resp, data)
	return data, err
}

// UserHasPicture checks if a user has a portrait.
func (c *Client) UserHasPicture(user, picture string, params *Parameters) (*Picture, error) {
	uri := fmt.Sprintf("/users/%s/pictures/%s", user, picture)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &Picture{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// UserEditPicture edits a portrait
func (c *Client) UserEditPicture(user, picture string, params *Parameters) (*Picture, error) {
	uri := fmt.Sprintf("/users/%s/pictures/%s", user, picture)
	resp, err := c.Patch(uri, params)

	if err != nil {
		return nil, err
	}

	data := &Picture{}

	err = c.processRequestData(204, resp, data)
	return data, err
}

// DeleteUserPicture removes a portrait from your portrait list.
// This method requires a token with the "delete" scope.
func (c *Client) DeleteUserPicture(user, picture string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/pictures/%s", user, picture)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

//GetUserPorfolios gets a list of Portfolios created by a user.
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
func (c *Client) GetUserPorfolios(user string, params *Parameters) (*PortfolioData, error) {
	uri := fmt.Sprintf("/users/%s/portfolios", user)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &PortfolioData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// GetUserPortfolio gets a portfolio
func (c *Client) GetUserPortfolio(user, portfolio string, params *Parameters) (*Portfolio, error) {
	uri := fmt.Sprintf("/users/%s/portfolios/%s", user, portfolio)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &Portfolio{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

//GetUserPortfolioVideos gets the videos in this Portfolio.
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
func (c *Client) GetUserPortfolioVideos(user, portfolio string, params *Parameters) (*VideoData, error) {
	uri := fmt.Sprintf("/users/%s/portfolios/%s/videos", user, portfolio)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &VideoData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// UserPortfoliosContaintsVideo checks if a Portfolio contains a video.
func (c *Client) UserPortfoliosContaintsVideo(user, portfolio, video string, params *Parameters) (bool, error) {
	uri := fmt.Sprintf("/users/%s/portfolios/%s/videos/%s", user, portfolio, video)
	resp, err := c.Get(uri, params)

	if err != nil {
		return false, err
	}

	if resp.StatusCode == 204 {
		return true, nil
	} else if resp.StatusCode == 404 {
		return false, nil
	}
	c.log.Printf("Request returned %s.\n", resp.Status)
	return false, fmt.Errorf("request returned %s", resp.Status)
}

// UserPortfoliosAddVideo adds a video to the Portfolio.
// This method requires a token with the "edit" scope.
func (c *Client) UserPortfoliosAddVideo(user, portfolio, video string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/portfolios/%s/videos/%s", user, portfolio, video)
	resp, err := c.Put(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

// UserPortfoliosDeleteVideos deletes a video from the Portfolio.
// This method requires a token with the "interact" scope.
func (c *Client) UserPortfoliosDeleteVideos(user, portfolio, video string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/portfolios/%s/videos/%s", user, portfolio, video)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

// GetUserWatchedVideos removes a video from the Portfolio.
// This method requires a token with the "edit" scope.
func (c *Client) GetUserWatchedVideos(user string, params *Parameters) (*VideoData, error) {
	uri := fmt.Sprintf("/users/%s/watched/videos", user)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &VideoData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// DeleteUserWatchedVideos gets a list of watched videos
func (c *Client) DeleteUserWatchedVideos(user string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/watched/videos", user)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err

}

// DeleteUserWatchedVideo deletes all watched videos
func (c *Client) DeleteUserWatchedVideo(user, video string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/watched/videos/%s", user, video)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

// UserUploadVideo begins the video upload process.
// This method requires a token with the "upload" scope.
func (c *Client) UserUploadVideo(user string, params *Parameters) (*AlbumDataElement, error) {
	uri := fmt.Sprintf("/users/%s/videos", user)
	resp, err := c.Post(uri, params)

	if err != nil {
		return nil, err
	}

	data := &AlbumDataElement{}

	err = c.processRequestData(201, resp, data)
	return data, err
}

// GetUserVideos gets a list of videos uploaded by a user.
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
func (c *Client) GetUserVideos(user, params *Parameters) (*VideoData, error) {
	uri := fmt.Sprintf("%s/videos", user)
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

//GetUserVideo checks if a user owns a clip.
func (c *Client) GetUserVideo(user, video string, params *Parameters) (*VideoDataElement, error) {
	uri := fmt.Sprintf("/users/%s/videos/%s", user, video)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &VideoDataElement{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

//GetUserWatchLater gets the authenticated user's Watch Later queue.
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
func (c *Client) GetUserWatchLater(user string, params *Parameters) (*VideoDataElement, error) {
	uri := fmt.Sprintf("/users/%s/watchlater", user)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &VideoDataElement{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

//CheckVideoInUserWatchedlater checks if a video is in the authenticated user's Watch Later queue.
func (c *Client) CheckVideoInUserWatchedlater(user, video string, params *Parameters) (bool, error) {
	uri := fmt.Sprintf("/users/%s/watchlater/%s", user, video)
	resp, err := c.Get(uri, params)

	if err != nil {
		return false, err
	}

	if resp.StatusCode == 204 {
		return true, nil
	} else if resp.StatusCode == 404 {
		return false, nil
	}
	c.log.Printf("Request returned %s.\n", resp.Status)
	return false, fmt.Errorf("request returned %s", resp.Status)
}

// UserAddWatchLater adds a video to the authenticated user's watch later list.
// This method requires a token with the "interact" scope.
func (c *Client) UserAddWatchLater(user, video string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/watchlater/%s", user, video)
	resp, err := c.Put(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

// DeleteUserWatchlaterVideo removes a video from your watch later list.
// This method requires a token with the "interact" scope.
func (c *Client) DeleteUserWatchlaterVideo(user, video string, params *Parameters) error {
	uri := fmt.Sprintf("/users/%s/watchlater/%s", user, video)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}
