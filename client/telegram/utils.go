package telegram

import "path"

func makeBaseUrl(token string) string {
	return "bot" + token
}

func (c *Client) makePath(method string) string {
	return path.Join(c.baseUrl, method)
}
