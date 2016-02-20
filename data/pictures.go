package data

type PicturesData struct {
	Total   int       `json:"total,omitempty"`
	Page    int       `json:"page, omitempty"`
	PerPage int       `json:"per_page, omitempty"`
	Paging  Paging    `json:"paging, omitempty"`
	Data    []Picture `json:"data,omitempty"`
}
