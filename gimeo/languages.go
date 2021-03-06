package gimeo

import (
	. "github.com/julianedialkova/gimeo/data"
)

// GetLanguages lists all valid video languages
func (c *Client) GetLanguages(params *Parameters) (*CommonsData, error) {
	resp, err := c.Get("/languages", params)

	if err != nil {
		return nil, err
	}

	data := &CommonsData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}
