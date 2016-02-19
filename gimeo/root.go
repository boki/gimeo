package gimeo

import (
	"fmt"

	. "github.com/julianedialkova/gimeo/data"
)

//List all endpoints
func (c *Client) ListAll() (*Endpoints, error) {
	resp, err := c.Get("/", nil)

	if err != nil {
		fmt.Println("ListAll: Could not execute request")
		return nil, err
	}

	data := &Endpoints{}

	err = c.processRequestData(200,resp, data)
	return data, err

}
