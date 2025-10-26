package telegram

import (
	"encoding/json"
	"net/url"
	"strconv"
)

func (c *Client) Updates(offset int) ([]UpdateEntity, error) {

	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))

	resp, err := c.doRequest(reqParams{
		apiMethod: getUpdatesMethod,
		query:     q,
	})

	if err != nil {
		return nil, err
	}

	var respData UpdatesResponse

	if err := json.Unmarshal(resp, &respData); err != nil {
		return nil, err
	}

	return respData.Result, nil

}
