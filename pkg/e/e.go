package e

import (
	"errors"
	"tg-bot/pkg/messages"
)

var ErrClientError = errors.New("")
var ErrServerError = errors.New("")

const ErrServerMSG = `
Внутренняя ошибка. 

Просим прощения. Повторите ваш запрос
`

var ErrAnswerTimeout = errors.Join(ErrClientError, errors.New("timeout session closed"))
var ErrSessionClosed = errors.New("session closed")
var ErrGroupNotFound = errors.Join(ErrClientError, errors.New(messages.TaskGroupNotValid))
