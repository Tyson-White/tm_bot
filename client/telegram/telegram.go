package telegram

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	protocol string
	host     string
	baseUrl  string
	client   http.Client
}

const (
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
)

func New(token string) Client {

	return Client{
		protocol: "https",
		host:     "api.telegram.org",
		baseUrl:  makeBaseUrl(token),
		client:   http.Client{},
	}
}

func (c *Client) Updates(offset int) ([]UpdateEntity, error) {

	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))

	resp, err := c.doRequest(getUpdatesMethod, q, nil)

	if err != nil {
		return nil, err
	}

	var respData UpdatesResponse

	if err := json.Unmarshal(resp, &respData); err != nil {
		return nil, err
	}

	return respData.Result, nil

}

func (c *Client) doRequest(apiMethod string, query url.Values, body any) ([]byte, error) {

	url := &url.URL{
		Scheme:   c.protocol,
		Host:     c.host,
		Path:     c.makePath(apiMethod),
		RawQuery: query.Encode(),
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	rBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return rBody, nil

}

func (c *Client) SendMessage(chatId, msg string) (MessageEntity, error) {
	q := url.Values{}

	q.Add("chat_id", chatId)
	q.Add("text", msg)

	resp, err := c.doRequest(sendMessageMethod, q, nil)

	if err != nil {
		return MessageEntity{}, err
	}

	var data MessageEntity

	err = json.Unmarshal(resp, &data)

	if err != nil {
		return MessageEntity{}, err
	}

	return data, nil
}
