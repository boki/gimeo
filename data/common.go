package data

type Picture struct {
	URI    string `json:"uri,omitempty"`
	Active bool   `json:"active,omitempty"`
	Type   string `json:"type,omitempty"`
	Sizes  []Size `json:"sizes,omitempty"`
}

type Size struct {
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
	Link   string `json:"link,omitempty"`
}

type InteractionInfo struct {
	Added     bool   `json:"added,omitempty"`
	AddedTime string `json:"added_time,omitempty"`
	Type      string `json:"type,omitempty"`
	Title     string `json:"title,omitempty"`
	URI       string `json:"uri,omitempty"`
}

type ConnectionInfo struct {
	Path    string   `json:"uri,omitempty"`
	Methods []string `json:"options,omitempty"`
	Total   int      `json:"total,omitempty"`
}

type Channels struct {
	Total int `json"total,omitempty"`
}

type Paging struct {
	Next     string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
	First    string `json:"first,omitempty"`
	Last     string `json:"last,omitempty"`
}
