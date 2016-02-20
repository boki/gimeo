package data

type VideoData struct {
	Total   int                `json:"total,omitempty"`
	Page    int                `json:"page, omitempty"`
	PerPage int                `json:"per_page, omitempty"`
	Paging  Paging             `json:"paging, omitempty"`
	Data    []VideoDataElement `json:"data,omitempty"`
}

type VideoDataElement struct {
	URI           string        `json:"uri,omitempty"`
	Name          string        `json:"name,omitempty"`
	Description   string        `json:"description"`
	Link          string        `json:"link,omitempty"`
	Duration      int           `json:"duration,omitempty"`
	Width         int           `json:"width,omitempty"`
	Language      string        `json:"language,omitempty"`
	Height        int           `json:"height,omitempty"`
	Embed         Embed         `json:"embed,omitempty"`
	ContentRating []string      `json:"content_rating"`
	License       string        `json:"license"`
	CreatedTime   string        `json:"created_time,omitempty"`
	ModifiedTime  string        `json:"modified_time,omitemtpy"`
	User          User          `json:"user, omitempty"`
	Pictures      Picture       `json:"pictures,omitempty"`
	Metadata      VideoMetadata `json:"metadata,omitempty"`
	Privacy       VideoPrivacy  `json:"privacy,omitempty"`
	Tags          []Tag         `json:"tags,omitempty"`
	Stats         Stats         `json:"stats,omitempty"`
	App           App           `json:"app, omitempty"`
	Status        string        `json:"status, omitempty"`
	// EmbedPresets  ?        `json:"embed_presets, omitempty"`
}

type App struct {
	Name string `json:"uri, omitempty"`
	URI  string `json:"name, omitempty"`
}

type Stats struct {
	Plays int `json:"plays,omitempty"`
}

type Embed struct {
	HTML string `json:"html,omitempty"`
}

type Tag struct {
	URI       string       `json:"uri,omitempty"`
	Name      string       `json:"name,omitempty"`
	Tag       string       `json:"tag,omitempty"`
	Canonical string       `json:"canonical,omitempty"`
	Metadata  TagsMetadata `json:"metadata,omitempty"`
}

type TagsMetadata struct {
	TagsConnections TagsConnections `json:"connections,omitempty"`
}

type TagsConnections struct {
	Videos ConnectionInfo `json:"videos,omitempty"`
}

type VideoPrivacy struct {
	View     string `json:"view,omitempty"`
	Embed    string `json:"embed,omitempty"`
	Download bool   `json:"download,omitempty"`
	Add      bool   `json:"add,omitempty"`
	Comments string `json:"comments,omitempty"`
}

type VideoMetadata struct {
	VideoConnections VideoConnections `json:"connections,omitempty"`
	VideoInteractios VideoInteractios `json:"interactions,omitempty"`
}

type VideoConnections struct {
	Likes      ConnectionInfo `json:"likes,omitempty"`
	Pictures   ConnectionInfo `json:"pictures,omitempty"`
	Comments   ConnectionInfo `json:"comments,omitempty"`
	Credits    ConnectionInfo `json:"credits,omitempty"`
	Texttracks ConnectionInfo `json:"texttracks,omitempty"`
	Related    ConnectionInfo `json:"related,omitempty"`
}

type VideoInteractios struct {
	WatchLater InteractionInfo `json:"watchlater,omitempty"`
	Like       InteractionInfo `json:"like,omitempty"`
}
