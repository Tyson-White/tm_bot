package telegram

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
)

const (
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
	sendPhotoMethod   = "sendPhoto"
)

func New(token string) Client {

	return Client{
		protocol: "https",
		host:     "api.telegram.org",
		baseUrl:  makeBaseUrl(token),
		client:   http.Client{},
	}
}

type reqParams struct {
	contentType string
	httpMethod  string
	apiMethod   string
	query       url.Values
	body        []byte
}

func (c *Client) doRequest(params reqParams) ([]byte, error) {

	url := &url.URL{
		Scheme: c.protocol,
		Host:   c.host,
		Path:   c.makePath(params.apiMethod),
	}

	if params.query.Encode() != "" {
		url.RawQuery = params.query.Encode()
	}

	var req *http.Request

	if params.httpMethod == http.MethodPost {
		req, _ = http.NewRequest(params.httpMethod, url.String(), bytes.NewBuffer(params.body))
	} else {
		req, _ = http.NewRequest(params.httpMethod, url.String(), nil)
	}

	if params.contentType == "" {
		req.Header.Set("Content-Type", "application/json")
	} else {
		req.Header.Set("Content-Type", params.contentType)
	}

	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	rBody, err := io.ReadAll(resp.Body)
	//
	//if params.apiMethod == sendPhotoMethod {
	//	log.Println(params.apiMethod, string(rBody))
	//}

	if err != nil {
		return nil, err
	}

	return rBody, nil

}
