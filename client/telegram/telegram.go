package telegram

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

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

func (c *Client) doRequest(httpMethod string, apiMethod string, query url.Values, body any) ([]byte, error) {

	url := &url.URL{
		Scheme:   c.protocol,
		Host:     c.host,
		Path:     c.makePath(apiMethod),
		RawQuery: query.Encode(),
	}

	b, _ := json.Marshal(body)

	req, _ := http.NewRequest(httpMethod, url.String(), bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	rBody, err := io.ReadAll(resp.Body)

	log.Println(string(rBody))

	if err != nil {
		return nil, err
	}

	return rBody, nil

}

func (c *Client) Updates(offset int) ([]UpdateEntity, error) {

	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))

	resp, err := c.doRequest(http.MethodGet, getUpdatesMethod, q, nil)

	if err != nil {
		return nil, err
	}

	var respData UpdatesResponse

	if err := json.Unmarshal(resp, &respData); err != nil {
		return nil, err
	}

	return respData.Result, nil

}

func (c *Client) SendMessage(chatId, msg string) (MessageEntity, error) {
	q := url.Values{}

	q.Add("chat_id", chatId)
	q.Add("text", msg)

	resp, err := c.doRequest(http.MethodGet, sendMessageMethod, q, nil)

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

func (c *Client) SendFMessage(chatId, msg string) (MessageEntity, error) {

	resp, err := c.doRequest(http.MethodPost, sendMessageMethod, nil, SendMessageBody{
		Chat: chatId,
		Text: msg,
		Mode: "HTML",
	})

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
