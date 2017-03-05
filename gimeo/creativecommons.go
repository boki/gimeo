package gimeo

import (
	. "github.com/julianedialkova/gimeo/data"
)

//GetCreativeCommons gets all valid creative commons licenses
func (c *Client) GetCreativeCommons(params *Parameters) (*CommonsData, error) {
	resp, err := c.Get("/creativecommons", params)

	if err != nil {
		return nil, err
	}

	data := &CommonsData{}

	err = c.processRequestData(200, resp, data)
	return data, err
}
