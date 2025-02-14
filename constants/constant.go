package constants

import "regexp"

const (
	HeaderContentType   = "Content-Type"
	HeaderAccept        = "accept"
	MIMEApplicationJSON = "application/json"

	MessageOk      = "ok"
	MessageSuccess = "SUCCESS"

	ErrBadRequest          = "Bad request received from client"
	ErrInvalidMethod       = "Method not allowed"
	ErrInternalServerError = "An error occured on the server. Please try again later."

	CtxTodoId string = "todoId"
)

var (
	RegexTodoId = regexp.MustCompile("/^[0-9A-F]{8}-[0-9A-F]{4}-1[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$/i")
)
