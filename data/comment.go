package data

//CommentData is the basic type, returned when working with comments.
type CommentData struct {
	Total   int                  `json:"total,omitempty"`
	Page    int                  `json:"page, omitempty"`
	PerPage int                  `json:"per_page, omitempty"`
	Paging  Paging               `json:"paging, omitempty"`
	Data    []CommentDataElement `json:"data,omitempty"`
}

//CommentDataElement is the basic type, returned when working with a comment.
type CommentDataElement struct {
	URI       string          `json:"uri,omitempty"`
	Type      string          `json:"type,omitempty"`
	Text      string          `json:"text,omitempty"`
	CreatedOn string          `json:"created_on,omitempty"`
	User      User            `json:"user, omitempty"`
	Metadata  CommentMetadata `json:"metadata,omitempty"`
}

//CommentMetadata is the basic type, where the comments store their metadata.
type CommentMetadata struct {
	CommentConnections CommentConnections `json:"connections,omitempty"`
}

//CommentConnections stores the basic connections from the Metadata.
type CommentConnections struct {
	Replies ConnectionInfo `json:"replies,omitempty"`
}
