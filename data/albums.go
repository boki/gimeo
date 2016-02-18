package data

type AlbumData struct {
	Total    int                `json:"total,omitempty"`
	Page     int                `json:"page, omitempty"`
	Per_page int                `json:"per_page, omitempty"`
	Paging   Paging             `json:"paging, omitempty"`
	Data     []AlbumDataElement `json:"data,omitempty"`
}

type AlbumDataElement struct {
	URI           string        `json:"uri,omitempty"`
	Name          string        `json:"name,omitempty"`
	Description   string        `json:"description"`
	Link          string        `json:"link,omitempty"`
	Duration      int           `json:"duration,omitempty"`
	CreatedTime   string        `json:"created_time,omitempty"`
	ModifiedTime  string        `json:"modified_time,omitemtpy"`
	User          User          `json:"user, omitempty"`
	Pictures      Picture       `json:"pictures,omitempty"`
	Metadata      AlbumMetadata `json:"metadata,omitempty"`
	Privacy       AlbumPrivacy  `json:"privacy,omitempty"`
}

type AlbumPrivacy struct {
	View     string `json:"view,omitempty"`
}

type AlbumMetadata struct {
	AlbumConnections AlbumConnections `json:"connections,omitempty"`
// 	AlbumInteractios AlbumInteractios `json:"interactions,omitempty"`
}

type AlbumConnections struct {
	Videos ConnectionInfo `json:"videos,omitempty"`
}

// type AlbumInteractios struct {
// 	WatchLater InteractionInfo `json:"watchlater,omitempty"`
// 	Like InteractionInfo `json:"like,omitempty"`
	
// }