package gimeo

import (
	. "github.com/julianedialkova/gimeo/data"
)

// ListAll lists all endpoints as URI templates
func (c *Client) ListAll() (*Endpoints, error) {
	resp, err := c.Get("/", nil)

	if err != nil {
		return nil, err
	}

	data := &Endpoints{}

	err = c.processRequestData(200, resp, data)
	return data, err

}
