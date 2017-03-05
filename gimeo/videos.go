package gimeo

import (
	"fmt"

	. "github.com/julianedialkova/gimeo/data"
)

// GetVideos searches for videos
// Name					Type					Required						Description
// page					int						No									The page number to show.
// per_page			int						No									Number of items to show on each page. Max 50.
// query				string				Yes									Search query.
// sort					string				No									Technique used to sort the results.
// 							relevant
// 							date
// 							alphabetical
// 							plays
// 							likes
// 							comments
// 							duration
// direction		string				No									The direction that the results are sorted.
// 							asc
// 							desc
// filter				string				No									Filter to apply to the results. The CC filters will show only those videos with the applicable creative commons licenses. See our Creative Commons page for more.
// 							CC
// 							CC-BY
// 							CC-BY-SA
// 							CC-BY-ND
// 							CC-BY-NC
// 							CC-BY-NC-SA
// 							CC-BY-NC-ND
// 							in-progress
func (c *Client) GetVideos(params *Parameters) (*VideoData, error) {
	resp, err := c.Get("/videos", params)

	if err != nil {
		return nil, err
	}

	data := &VideoData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// GetVideo gets a video.
func (c *Client) GetVideo(video string, params *Parameters) (*VideoDataElement, error) {
	uri := fmt.Sprintf("/videos/%s", video)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &VideoDataElement{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// ChangeVideo edits video metadata.
// This method requires a token with the "edit" scope.
// Name												Type				Required			Description
// name												string			No						The new title for the video
// description								string			No						The new description for the video
// license										string			No						Set the Creative Commons license
// 														by
// 														by-sa
// 														by-nd
// 														by-nc
// 														by-nc-sa
// 														by-nc-nd
// 														cc0
// privacy.view								string			No					The new privacy setting for the video. Content-type application/json is the only valid type for type "users", basic users can not set privacy to unlisted.
// 														anybody
// 														nobody
// 														contacts
// 														password
// 														users
// 														unlisted
// 														disable
// privacy.download						boolean			No					Enable or disable the ability for anyone to download video.
// privacy.add	boolean										No					Enable or disable the ability for anyone to add the video to an album, channel, or group.
// privacy.comments						string			No					The privacy for who can comment on the video.
// 														anybody
// 														nobody
// 														contacts
// password										string			No					When you set privacy.view to password, you must provide the password as an additional parameter
// privacy.embed							string			No					The videos new embed settings. Whitelist allows you to define all valid embed domains. Check out our docs for adding and removing domains.
// 														public
// 														private
// 														whitelist
// review_link								boolean			No					Enable or disable the review page
// locale											string			No					Set the default language for this video. For a full list of valid languages use the "/languages?filter=texttracks" endpoint
// content_rating							array				No					A list of values describing the content in this video. You can find the full list in the /contentrating endpoint. You must provide a list representation appropriate for your request body (comma separated for querystring, or array for JSON)
// embed.buttons.like					boolean			No					Show or hide the like button
// embed.buttons.watchlater		boolean			No					Show or hide the watch later button
// embed.buttons.share				boolean			No					Show or hide the share button
// embed.buttons.embed				boolean			No					Show or hide the embed button
// embed.buttons.hd						boolean			No					Show or hide the hd button
// embed.buttons.fullscreen		boolean			No					Show or hide the fullscreen button
// embed.buttons.scaling			boolean			No					Show or hide the scaling button (shown only in fullscreen mode)
// embed.logos.vimeo					boolean			No					Show or hide the vimeo logo
// embed.logos.custom.active	boolean			No					Show or hide your custom logo
// embed.logos.custom.stick		boolean			No					Always show the custom logo, or hide it after time with the rest of the UI
// embed.logos.custom.lin			string			No					A url that your user will navigate to if they click your custom logo
// embed.playbar							boolean			No					Show or hide the playbar
// embed.volume								boolean			No					Show or hide the volume selector
// embed.color								string			No					A primary color used by the embed player
// embed.title.owner					string			No					Show, hide, or let the user decide if the owners information shows on the video
// 														user
// 														show
// 														hide
// embed.title.portrait				string			No						Show, hide, or let the user decide if the owners portrait shows on the video
// 														user
// 														show
// 														hide
// embed.title.name						string			No						Show, hide, or let the user decide if the video title shows on the video
// 														user
// 														show
// 														hide
func (c *Client) ChangeVideo(video string, params *Parameters) (*VideoDataElement, error) {
	uri := fmt.Sprintf("/videos/%s", video)
	resp, err := c.Patch(uri, params)

	if err != nil {
		return nil, err
	}

	data := &VideoDataElement{}

	err = c.processRequestData(204, resp, data)
	return data, err
}

// DeleteVideo deletes a video
// This method requires a token with the "delete" scope.
func (c *Client) DeleteVideo(video string, params *Parameters) error {
	uri := fmt.Sprintf("/videos/%s", video)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

// ReplaceVideo gets an upload ticket to replace this video file.
// This method requires a token with the "upload" scope.
// Name							Type				Required				Description
// type							string			Yes							Upload type
// 									POST
// 									streaming
// redirect_url			string			Yes							The app redirect URL
func (c *Client) ReplaceVideo(video string, params *Parameters) error {
	uri := fmt.Sprintf("/videos/%s/files", video)
	resp, err := c.Put(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

//GetVideoCredits gets a list of users credited on a video.
// Name								Type				Required			Description
// page								int					No						The page number to show.
// per_page						int					No						Number of items to show on each page. Max 50.
// query							string			No						Search query.
// sort								string			No						Technique used to sort the results.
// 										date
// 										alphabetical
// direction					string			No						The direction that the results are sorted.
// 										asc
// 										desc
func (c *Client) GetVideoCredits(video string, params *Parameters) (*CreditsData, error) {
	uri := fmt.Sprintf("/videos/%s/credits", video)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &CreditsData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

//AddCreditToVideo adds a credit to a video.
// Name			Type		Required	Description
// role			string	Yes				The role of the person being credited
// name			string	Yes				The name of the person being credited
// email		string	Yes				The email address of the person being credited
// user_uri	string	Yes				The URI of the Vimeo user who should be given credit in this video
func (c *Client) AddCreditToVideo(video string, params *Parameters) (*CreditsDataElement, error) {
	uri := fmt.Sprintf("/videos/%s/credits", video)
	resp, err := c.Post(uri, params)

	if err != nil {
		return nil, err
	}

	data := &CreditsDataElement{}

	err = c.processRequestData(201, resp, data)
	return data, err
}

// GetRelatedVideos gets related videos.
// Name								Type				Required			Description
// page								int					No						The page number to show.
// per_page						int					No						Number of items to show on each page. Max 50.
// filter							string			Yes						Filter to apply to the results.
// 										related
func (c *Client) GetRelatedVideos(video string, params *Parameters) (*VideoData, error) {
	uri := fmt.Sprintf("/videos/%s/videos", video)
	resp, err := c.Get(uri, params)

	if err != nil {
		return nil, err
	}

	data := &VideoData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// GetVideoCategories gets a list of all categories this video is in.
// Name								Type				Required			Description
// page								int					No						The page number to show.
// per_page						int					No						Number of items to show on each page. Max 50.
func (c *Client) GetVideoCategories(video string, params *Parameters) (*CategoryData, error) {
	uri := fmt.Sprintf("/videos/%s/categories", video)
	resp, err := c.Get(uri, params)
	if err != nil {
		return nil, err
	}

	data := &CategoryData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// AddVideoCategories adds up to 2 categories and 1 sub-category to a video. This is mearly a suggestion, and does not ensure that the video will added to the category.
// This method requires a token with the "edit" scope.
func (c *Client) AddVideoCategories(video string, params *Parameters) error {
	uri := fmt.Sprintf("/videos/%s/categories", video)
	resp, err := c.Put(uri)

	if err != nil {
		return err
	}

	return c.processRequestNoData(resp)
}

// GetVideoCredit gets a single credit.
func (c *Client) GetVideoCredit(video, credit string, params *Parameters) (*CreditsDataElement, error) {
	uri := fmt.Sprintf("/videos/%s/credits/%s", video, credit)
	resp, err := c.Get(uri, params)
	if err != nil {
		return nil, err
	}

	data := &CreditsDataElement{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// ChangeVideoCredit edits information about a single credit
// This method requires a token with the "edit" scope.
func (c *Client) ChangeVideoCredit(video, credit string, params *Parameters) (*Picture, error) {
	uri := fmt.Sprintf("/videos/%s/credits/%s", video, credit)
	resp, err := c.Patch(uri, params)

	if err != nil {
		return nil, err
	}

	data := &Picture{}

	err = c.processRequestData(204, resp, data)
	return data, err
}

// DeleteVideoCredit deletes a credit
// This method requires a token with the "edit" scope.
func (c *Client) DeleteVideoCredit(video, credit string, params *Parameters) error {
	uri := fmt.Sprintf("/videos/%s/credits/%s", video, credit)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

//GetVideoComments gets comments on this video.
// Name								Type				Required			Description
// page								int					No						The page number to show.
// per_page						int					No						Number of items to show on each page. Max 50.
// query							string			No						Search query.
// sort								string			No						Technique used to sort the results.
// 										date
// 										alphabetical
// direction					string			No						The direction that the results are sorted.
// 										asc
// 										desc
func (c *Client) GetVideoComments(video string, params *Parameters) (*CommentData, error) {
	uri := fmt.Sprintf("/videos/%s/comments", video)
	resp, err := c.Get(uri, params)
	if err != nil {
		return nil, err
	}

	data := &CommentData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// CommentVideo posts a comment on the video
func (c *Client) CommentVideo(video string, params *Parameters) (*CommentDataElement, error) {
	uri := fmt.Sprintf("/videos/%s/comments", video)
	resp, err := c.Post(uri, params)

	if err != nil {
		return nil, err
	}

	data := &CommentDataElement{}

	err = c.processRequestData(201, resp, data)
	return data, err
}

// ChangeVideoComment edits an existing comment on a video
func (c *Client) ChangeVideoComment(video, comment string, params *Parameters) (*CommentDataElement, error) {
	uri := fmt.Sprintf("/videos/%s/comments/%s", video, comment)
	resp, err := c.Patch(uri, params)

	if err != nil {
		return nil, err
	}

	data := &CommentDataElement{}

	err = c.processRequestData(204, resp, data)
	return data, err
}

// DeleteVideoComment deletes a comment from a video
// This method requires a token with the "delete" scope.
func (c *Client) DeleteVideoComment(video, comment string, params *Parameters) error {
	uri := fmt.Sprintf("/videos/%s/comments/%s", video, comment)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

// GetVideoCommentReplies gets comments on this video.
// Name								Type				Required			Description
// page								int					No						The page number to show.
// per_page						int					No						Number of items to show on each page. Max 50.
func (c *Client) GetVideoCommentReplies(video, comment string, params *Parameters) (*CommentData, error) {
	uri := fmt.Sprintf("/videos/%s/comments/%s/replies", video, comment)
	resp, err := c.Get(uri, params)
	if err != nil {
		return nil, err
	}

	data := &CommentData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// ReplyToCommentVideo posts a reply to a comment on the video
// 	Name	Type		Required	Description
//	text	string	Yes				The comment's new text
func (c *Client) ReplyToCommentVideo(video, comment string, params *Parameters) (*CommentDataElement, error) {
	uri := fmt.Sprintf("/videos/%s/comments/%s/replies", video, comment)
	resp, err := c.Post(uri, params)

	if err != nil {
		return nil, err
	}

	data := &CommentDataElement{}

	err = c.processRequestData(201, resp, data)
	return data, err
}

// GetVideoPictures gets a list of this video's past and present pictures.
func (c *Client) GetVideoPictures(video string, params *Parameters) (*PicturesData, error) {
	uri := fmt.Sprintf("/videos/%s/pictures", video)
	resp, err := c.Get(uri, params)
	if err != nil {
		return nil, err
	}

	data := &PicturesData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// AddPictureToVideo adds a picture resource to a video
// This method requires a token with the "upload" scope.
func (c *Client) AddPictureToVideo(video string, params *Parameters) (*VideoDataElement, error) {
	uri := fmt.Sprintf("/videos/%s/pictures", video)
	resp, err := c.Post(uri, params)

	if err != nil {
		return nil, err
	}

	data := &VideoDataElement{}

	err = c.processRequestData(201, resp, data)
	return data, err
}

// GetVideoPicture gets a single picture resource for a video
func (c *Client) GetVideoPicture(video, picture string, params *Parameters) (*Picture, error) {
	uri := fmt.Sprintf("/videos/%s/pictures/%s", video, picture)
	resp, err := c.Get(uri, params)
	if err != nil {
		return nil, err
	}

	data := &Picture{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// ChangeVideoPicture modifies an existing picture on a video.
// This method requires a token with the "edit" scope.
func (c *Client) ChangeVideoPicture(video, picture string, params *Parameters) (*Picture, error) {
	uri := fmt.Sprintf("/videos/%s/pictures/%s", video, picture)
	resp, err := c.Patch(uri, params)

	if err != nil {
		return nil, err
	}

	data := &Picture{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// DeleteVideoPicture removes an existing picture from a video.
// This method requires a token with the "edit" scope.
func (c *Client) DeleteVideoPicture(video, picture string, params *Parameters) error {
	uri := fmt.Sprintf("/videos/%s/pictures/%s", video, picture)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

// GetVideoLikes gets a list of the users who liked this video.
// Name								Type				Required			Description
// page								int					No						The page number to show.
// per_page						int					No						Number of items to show on each page. Max 50.
// query							string			No						Search query.
// sort								string			No						Technique used to sort the results.
// 										date
// 										alphabetical
// direction					string			No						The direction that the results are sorted.
// 										asc
// 										desc
func (c *Client) GetVideoLikes(video string, params *Parameters) (*LikesData, error) {
	uri := fmt.Sprintf("/videos/%s/likes", video)
	resp, err := c.Get(uri, params)
	if err != nil {
		return nil, err
	}

	data := &LikesData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// GetVideoTags lists all of the tags on the video
func (c *Client) GetVideoTags(video string, params *Parameters) (*TagsData, error) {
	uri := fmt.Sprintf("/videos/%s/tags", video)
	resp, err := c.Get(uri, params)
	if err != nil {
		return nil, err
	}

	data := &TagsData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// CheckVideoTag checks if a tag has been applied to a video
func (c *Client) CheckVideoTag(video, tag string, params *Parameters) (*TagsData, error) {
	uri := fmt.Sprintf("/videos/%s/tags/%s", video, tag)
	resp, err := c.Get(uri, params)
	if err != nil {
		return nil, err
	}

	data := &TagsData{}

	err = c.processRequestData(204, resp, data)
	return data, err
}

// VideoPutTag tags a video.
// This method requires a token with the "edit" scope.
func (c *Client) VideoPutTag(video, tag string, params *Parameters) error {
	uri := fmt.Sprintf("/videos/%s/tags/%s", video, tag)
	resp, err := c.Put(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

// DeleteVideoTag removes the tag from this video
// This method requires a token with the "edit" scope.
func (c *Client) DeleteVideoTag(video, tag string, params *Parameters) error {
	uri := fmt.Sprintf("/videos/%s/tags/%s", video, tag)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

// GetVideoUsers gets all users that are allowed to see this video.
func (c *Client) GetVideoUsers(video string, params *Parameters) (*UserData, error) {
	uri := fmt.Sprintf("/videos/%s/privacy/users", video)
	resp, err := c.Get(uri, params)
	if err != nil {
		return nil, err
	}

	data := &UserData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// AddVideoUsers adds users to the allowed users list.
// This method requires a token with the "edit" scope.
func (c *Client) AddVideoUsers(video string, params *Parameters) (*UserData, error) {
	uri := fmt.Sprintf("/videos/%s/privacy/users", video)
	resp, err := c.Put(uri)

	if err != nil {
		return nil, err
	}
	data := &UserData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}

// AddVideoUser adds a user to the allowed users list.
// This method requires a token with the "edit" scope.
func (c *Client) AddVideoUser(video, user string, params *Parameters) error {
	uri := fmt.Sprintf("/videos/%s/privacy/users/%s", video, user)
	resp, err := c.Put(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}

//DeleteVideoUser removes a user from the allowed users list.
// his method requires a token with the "edit" scope.
func (c *Client) DeleteVideoUser(video, user string, params *Parameters) error {
	uri := fmt.Sprintf("/videos/%s/privacy/users/%s", video, user)
	resp, err := c.Delete(uri)

	if err != nil {
		return err
	}

	err = c.processRequestNoData(resp)
	return err
}
