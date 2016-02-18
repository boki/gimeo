package data

type CommonsData struct {
	Total    int                `json:"total,omitempty"`
	Page     int                `json:"page, omitempty"`
	Per_page int                `json:"per_page, omitempty"`
	Paging   Paging             `json:"paging, omitempty"`
	Data     []CommonsDataElement `json:"data,omitempty"`
}

type CommonsDataElement struct {
    Code string `json:"code,omitempty"`
    Name string `json:"name,omitempty"`

}