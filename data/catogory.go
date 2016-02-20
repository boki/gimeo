package data

//CategoryData is the basic type, returned when working with a category.
type CategoryData struct {
	Total   int                   `json:"total,omitempty"`
	Page    int                   `json:"page, omitempty"`
	PerPage int                   `json:"per_page, omitempty"`
	Paging  Paging                `json:"paging, omitempty"`
	Data    []CategoryDataElement `json:"data,omitempty"`
}

//CategoryDataElement is the basic type, returned when working with a category.
type CategoryDataElement struct {
	URI                   string           `json:"uri,omitempty"`
	Name                  string           `json:"name,omitempty"`
	Link                  string           `json:"link,omitempty"`
	TopLevel              bool             `json:"top_level,omitempty"`
	Pictures              Picture          `json:"pictures,omitempty"`
	LastVideoFeaturedTime string           `json:"last_video_featured_time,omitempty"`
	Parent                string           `json:"parent,omitempty"`
	Metadata              CategoryMetadata `json:"metadata,omitempty"`
	Subcategories         []Subcategory    `json:"subcategories,omitempty"`
}

//CategoryMetadata is the basic type, where the categories store their metadata.
type CategoryMetadata struct {
	CategoryConnections CategoryConnections `json:"connections,omitempty"`
}

//CategoryConnections stores the basic connections from the Metadata.
type CategoryConnections struct {
	Channels ConnectionInfo `json:"channels,omitempty"`
	Groups   ConnectionInfo `json:"groups,omitempty"`
	Users    ConnectionInfo `json:"users,omitempty"`
	Videos   ConnectionInfo `json:"videos,omitempty"`
}

// Subcategory stores the subcategory information.
type Subcategory struct {
	URI  string `json:"uri,omitempty"`
	Name string `json:"name,omitempty"`
	Link string `json:"link,omitempty"`
}
