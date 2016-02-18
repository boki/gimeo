package data

type ChannelData struct {
	Total    int                `json:"total,omitempty"`
	Page     int                `json:"page, omitempty"`
	Per_page int                `json:"per_page, omitempty"`
	Paging   Paging             `json:"paging, omitempty"`
	Data     []ChannelDataElement `json:"data,omitempty"`
}

type ChannelDataElement struct {
	URI          string        `json:"uri,omitempty"`
	Name         string        `json:"name,omitempty"`
	Description  string        `json:"description"`
	Link         string        `json:"link,omitempty"`
	CreatedTime  string        `json:"created_time,omitempty"`
	ModifiedTime string        `json:"modified_time,omitemtpy"`
	User         User          `json:"user, omitempty"`
	Pictures     Picture       `json:"pictures,omitempty"`
	Metadata     ChannelMetadata `json:"metadata,omitempty"`
	Privacy      ChannelPrivacy  `json:"privacy,omitempty"`
	Header       Picture       `json:"header, omitempty"`
}

type ChannelPrivacy struct {
	View    string `json:"view,omitempty"`
}

type ChannelConnections struct {
	Users  ConnectionInfo `json:"users,omitempty"`
	Videos ConnectionInfo `json:"videos,omitempty"`
}

type ChannelInteractions struct {
	Follow InteractionInfo `json:"follow,omitempty"`
}

type ChannelMetadata struct {
	ChannelConnections  ChannelConnections  `json:"connections,omitempty"`
	ChannelInteractions ChannelInteractions `json:"interactions,omitempty"`
}
