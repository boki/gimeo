package data

type UserData struct {
	Total    int                `json:"total,omitempty"`
	Page     int                `json:"page, omitempty"`
	Per_page int                `json:"per_page, omitempty"`
	Paging   Paging             `json:"paging, omitempty"`
	Data     []User `json:"data,omitempty"`
}

type User struct {
	URI         string       `json:"uri,omitempty"`
	Name        string       `json:"name,omitempty"`
	Link        string       `json:"link,omitempty"`
	Location    string       `json:"location,omitempty"`
	Bio         string       `json:"bio,omitempty"`
	CreatedTime string       `json:"created_time,omitempty"`
	Account     string       `json:"account,omitempty"`
	Pictures    Picture      `json:"pictures,omitempty"`
	Websites    []Websites   `json:"websites,omitempty"`
	Metadata    UserMetadata `json:"metadata,omitempty"`
	Preferences Preferences  `json:"preferences,omitempty"`
	ContentFilter []string `json:"content_filter,omitempty"`
}

type Websites struct {
	Name        string `json:"name,omitempty"`
	Link        string `json: "link,omitempty"`
	Description string `json:"description,omitempty"`
}


type UserMetadata struct {
	UserConnections  UserConnections  `json:"connections,omitempty"`
	UserInteractions UserInteractions `json:"interactions,omitempty"`
}

type UserConnections struct {
	Activities ConnectionInfo `json:"activities,omitempty"`
	Albums     ConnectionInfo `json:"albums,omitempty"`
	Feed       ConnectionInfo `json:"feed,omitempty"`
	Followers  ConnectionInfo `json:"followers,omitempty"`
	Following  ConnectionInfo `json:"following,omitempty"`
	Likes      ConnectionInfo `json:"likes,omitempty"`
	PortFolios ConnectionInfo `json:"portfolios,omitempty"`
	Shared     ConnectionInfo `json:"shared,omitempty"`
	Pictures   ConnectionInfo `json:"pictures,omitempty"`
	Channels   ConnectionInfo `json:"channels,omitempty"`
	Groups     ConnectionInfo `json:"groups,omitempty"`
	WatchLater ConnectionInfo `json:"watchlater,omitempty"`
	Videos     ConnectionInfo `json:"videos,omitempty"`
}

type UserInteractions struct {
	Follow InteractionInfo `json:"follow,omitempty"`
}

type Preferences struct {
	Videos Preference `json:"videos,omitempty"`
}

type Preference struct {
	Privacy string `json:"privacy,omitempty"`
}