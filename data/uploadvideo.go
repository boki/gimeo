package data

type Upload struct {
	URI              string `json:"uri,omitempty"`
	TicketID         string `json:"ticket_id,omitempty"`
	User             User   `json:"user, omitempty"`
	UploadLink       string `json:"upload_link, omitempty"`
	UploadLinkSecure string `json:"upload_link_secure, omitempty"`
	CompleteURI      string `json:"complete_uri, omitempty"`
}
