package data

type Oauth struct {
    AccessToken string `json:"access_token,omitempty"`
    TokenType string `json:"token_type,omitempty"`
    Scope string `json:"scope,omitempty"`
    App App `json:"app,omitempty"`
    User User `json:"user,omitempty"`
    
}