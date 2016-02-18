package data

type PortfolioData struct {
	Total    int                `json:"total,omitempty"`
	Page     int                `json:"page, omitempty"`
	Per_page int                `json:"per_page, omitempty"`
	Paging   Paging             `json:"paging, omitempty"`
	Data     []Portfolio `json:"data,omitempty"`
}

type Portfolio struct {
    Uri    string                `json:"uri,omitempty"`
    Name    string               `json:"name,omitempty"`
    Link    string               `json:"link,omitempty"`
    CretedTime    string         `json:"created_time,omitempty"`
    ModifiedTime string `json:"modified_time,omitempty"`
    Sort string `json:"sort,omitempty"`
    Metadata PortfolioMetadata `json:"metadata,omitempty"`
    
}

type PortfolioConnections struct {
	Videos ConnectionInfo `json:"videos,omitempty"`
}


type PortfolioMetadata struct {
	PortfolioConnections  PortfolioConnections  `json:"connections,omitempty"`
}