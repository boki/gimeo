package data

type LikesData struct {
	Total   int                `json:"total,omitempty"`
	Page    int                `json:"page, omitempty"`
	PerPage int                `json:"per_page, omitempty"`
	Paging  Paging             `json:"paging, omitempty"`
	Data    []VideoDataElement `json:"data,omitempty"`
}
