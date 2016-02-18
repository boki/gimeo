package data

type FeedData struct {
	Total    int                `json:"total,omitempty"`
	Page     int                `json:"page, omitempty"`
	Per_page int                `json:"per_page, omitempty"`
	Paging   Paging             `json:"paging, omitempty"`
	Data     []FeedDataElement `json:"data,omitempty"`
}

type FeedDataElement struct {
	URI          string        `json:"uri,omitempty"`
	Clip VideoDataElement `json:"clip,omitempty"`
	Type string `json:"type,omitempty"`
	Time string `json:"time,omitempty"`
	Metadata FeedMetadata `json:"clip,omitempty"`
	Channel ChannelDataElement `json:"channel"`
}


type FeedMetadata struct {
	FeedConnections  FeedConnections  `json:"connections,omitempty"`
}

type FeedConnections struct {
	Related ConnectionInfo `json:"related,omitempty"`
}
