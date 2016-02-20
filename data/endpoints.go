package data

type Endpoints struct {
	Endpoints []Endpoint `json:"endpoints,omitempty"`
}

type Endpoint struct {
	Path    string   `json:"path,omitempty"`
	Methods []string `json:"methods,omitempty"`
}
