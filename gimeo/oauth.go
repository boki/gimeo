package gimeo

import (
	. "github.com/julianedialkova/gimeo/data"
)

//VerifyToken verifies the users token
func (c *Client) VerifyToken(params *Parameters) (*Oauth, error) {
	resp, err := c.Get("/oauth/verify", params)

	if err != nil {
		return nil, err
	}

	data := &Oauth{}

	err = c.processRequestData(200, resp, data)
	return data, err
}
