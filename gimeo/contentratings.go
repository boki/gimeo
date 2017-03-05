package gimeo

import (
	. "github.com/julianedialkova/gimeo/data"
)

// GetContentRatings gets all valid content ratings
func (c *Client) GetContentRatings(params *Parameters) (*CommonsData, error) {
	resp, err := c.Get("/contentratings", params)

	if err != nil {
		return nil, err
	}

	data := &CommonsData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}
