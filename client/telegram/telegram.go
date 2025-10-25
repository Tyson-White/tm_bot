package telegram

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
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

	if params.apiMethod == sendPhotoMethod {
		log.Println(params.apiMethod, string(rBody))
	}

	if err != nil {
		return nil, err
	}

	return rBody, nil

}

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

func (c *Client) SendMessage(chatId, msg string) (MessageEntity, error) {
	q := url.Values{}

	q.Add("chat_id", chatId)
	q.Add("text", msg)

	resp, err := c.doRequest(reqParams{
		httpMethod: http.MethodPost,
		apiMethod:  sendMessageMethod,
		query:      q,
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

func (c *Client) SendFMessage(chatId, msg string) (MessageEntity, error) {

	body := SendMessageBody{
		Chat: chatId,
		Text: msg,
		Mode: "HTML",
	}

	b, _ := json.Marshal(body)
	resp, err := c.doRequest(reqParams{
		httpMethod: http.MethodPost,
		apiMethod:  sendMessageMethod,
		body:       b,
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

func (c *Client) SendPhoto(chatId, filepath, msg string) (MessageEntity, error) {
	var (
		buf = new(bytes.Buffer)
		wr  = multipart.NewWriter(buf)
	)

	wr.WriteField("chat_id", chatId)
	wr.WriteField("caption", msg)
	wr.WriteField("parse_mode", "HTML")

	file, _ := os.Open(filepath)
	defer file.Close()

	field, _ := wr.CreateFormFile("photo", "photo")
	io.Copy(field, file)

	wr.Close()

	// log.Println(buf.String())

	_, err := c.doRequest(reqParams{
		contentType: wr.FormDataContentType(),
		httpMethod:  http.MethodPost,
		apiMethod:   sendPhotoMethod,
		body:        buf.Bytes(),
	})

	return MessageEntity{}, err
}
