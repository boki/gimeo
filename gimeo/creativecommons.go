package gimeo

import (
	"fmt"

	. "github.com/julianedialkova/gimeo/data"
)

func (c *Client) GetCreativeCommons(params *Parameters) (*CommonsData, error) {
	resp, err := c.Get("/creativecommons", params)

	if err != nil {
		fmt.Println("Could not execute request")
		return nil, err
	}

	data := &CommonsData{}

	err = c.processRequestData(200,resp, data)
	return data, err
}
