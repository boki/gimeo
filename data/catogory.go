package data

type CategoryData struct {
	Total    int                   `json:"total,omitempty"`
	Page     int                   `json:"page, omitempty"`
	Per_page int                   `json:"per_page, omitempty"`
	Paging   Paging                `json:"paging, omitempty"`
	Data     []CategoryDataElement `json:"data,omitempty"`
}

type CategoryDataElement struct {
	URI  string `json:"uri,omitempty"`
	Name string `json:"name,omitempty"`
	// Description           string           `json:"description"`
	Link                  string           `json:"link,omitempty"`
	TopLevel              bool             `json:"top_level,omitempty"`
	Pictures              Picture          `json:"pictures,omitempty"`
	LastVideoFeaturedTime string           `json:"last_video_featured_time,omitempty"`
	Parent                string           `json:"parent,omitempty"`
	Metadata              CategoryMetadata `json:"metadata,omitempty"`
	Subcategories         []Subcategory    `json:"subcategories,omitempty"`
}

type CategoryMetadata struct {
	CategoryConnections CategoryConnections `json:"connections,omitempty"`
}

type CategoryConnections struct {
	Channels ConnectionInfo `json:"channels,omitempty"`
	Groups   ConnectionInfo `json:"groups,omitempty"`
	Users    ConnectionInfo `json:"users,omitempty"`
	Videos   ConnectionInfo `json:"videos,omitempty"`
}

type Subcategory struct {
	URI  string `json:"uri,omitempty"`
	Name string `json:"name,omitempty"`
	Link string `json:"link,omitempty"`
}
