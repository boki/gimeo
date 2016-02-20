package data

type GroupData struct {
	Total   int                `json:"total,omitempty"`
	Page    int                `json:"page, omitempty"`
	PerPage int                `json:"per_page, omitempty"`
	Paging  Paging             `json:"paging, omitempty"`
	Data    []GroupDataElement `json:"data,omitempty"`
}

type GroupDataElement struct {
	URI          string        `json:"uri,omitempty"`
	Name         string        `json:"name,omitempty"`
	Description  string        `json:"description"`
	Link         string        `json:"link,omitempty"`
	CreatedTime  string        `json:"created_time,omitempty"`
	ModifiedTime string        `json:"modified_time,omitemtpy"`
	User         User          `json:"user, omitempty"`
	Pictures     Picture       `json:"pictures,omitempty"`
	Metadata     GroupMetadata `json:"metadata,omitempty"`
	Privacy      GroupPrivacy  `json:"privacy,omitempty"`
	Header       Picture       `json:"header, omitempty"`
}

type GroupPrivacy struct {
	View    string `json:"view,omitempty"`
	Join    string `json:"join,omitempty"`
	Videos  string `json:"videos,omitempty"`
	Comment string `json:"comment,omitempty"`
	Forums  string `json:"forums,omitempty"`
	Invite  string `json:"invite,omitempty"`
}

type GroupMetadata struct {
	GroupConnections  GroupConnections  `json:"connections,omitempty"`
	GroupInteractions GroupInteractions `json:"interactions,omitempty"`
}

type GroupConnections struct {
	Users  ConnectionInfo `json:"users,omitempty"`
	Videos ConnectionInfo `json:"videos,omitempty"`
}

type GroupInteractions struct {
	Join InteractionInfo `json:"join,omitempty"`
}
