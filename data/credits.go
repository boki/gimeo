package data

type CreditsData struct {
	Total   int                  `json:"total,omitempty"`
	Page    int                  `json:"page, omitempty"`
	PerPage int                  `json:"per_page, omitempty"`
	Paging  Paging               `json:"paging, omitempty"`
	Data    []CreditsDataElement `json:"data,omitempty"`
}

type CreditsDataElement struct {
	URI   string           `json:"uri,omitempty"`
	Name  string           `json:"name,omitempty"`
	Role  string           `json:"role,omitempty"`
	Video VideoDataElement `json:"video,omitempty"`
	User  User             `json:"user,omitempty"`
}
