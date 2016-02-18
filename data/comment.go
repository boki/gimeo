package data

type CommentData struct {
	Total    int                `json:"total,omitempty"`
	Page     int                `json:"page, omitempty"`
	Per_page int                `json:"per_page, omitempty"`
	Paging   Paging             `json:"paging, omitempty"`
	Data     []CommentDataElement `json:"data,omitempty"`
}

type CommentDataElement struct {
	URI           string        `json:"uri,omitempty"`
	Type           string        `json:"type,omitempty"`
	Text string        `json:"text,omitempty"`
	CreatedOn   string        `json:"created_on,omitempty"`
	User          User          `json:"user, omitempty"`
	Metadata      CommentMetadata `json:"metadata,omitempty"`
}

type CommentMetadata struct {
	CommentConnections CommentConnections `json:"connections,omitempty"`
}

type CommentConnections struct {
	Replies      ConnectionInfo `json:"replies,omitempty"`
}
