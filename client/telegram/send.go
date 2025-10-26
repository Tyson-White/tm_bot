package telegram

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// Сообщение с форматированием
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

	resp, err := c.doRequest(reqParams{
		contentType: wr.FormDataContentType(),
		httpMethod:  http.MethodPost,
		apiMethod:   sendPhotoMethod,
		body:        buf.Bytes(),
	})

	var data MessageEntity

	err = json.Unmarshal(resp, &data)

	if err != nil {
		return MessageEntity{}, err
	}

	return data, err
}

func (c *Client) SendError(chatId, msg string) (MessageEntity, error) {
	return c.SendPhoto(chatId, "./assets/error.png", msg)
}

func (c *Client) SendSuccess(chatId, msg string) (MessageEntity, error) {
	return c.SendPhoto(chatId, "./assets/success.png", msg)
}
