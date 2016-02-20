package data

type FollowersData struct {
	Total   int    `json:"total,omitempty"`
	Page    int    `json:"page, omitempty"`
	PerPage int    `json:"per_page, omitempty"`
	Paging  Paging `json:"paging, omitempty"`
	Data    []User `json:"data,omitempty"`
}
