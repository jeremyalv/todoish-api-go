package constants

const (
	HeaderContentType   = "Content-Type"
	HeaderAccept        = "accept"
	MIMEApplicationJSON = "application/json"

	MessageOk      = "ok"
	MessageSuccess = "SUCCESS"

	ErrBadRequest          = "Bad request received from client"
	ErrPreconditionFailed  = "Precondition failed"
	ErrInvalidMethod       = "Method not allowed"
	ErrInternalServerError = "An error occured on the server. Please try again later."

	CtxTodoId string = "todoId"

	TodoEndpoint = "/api/{version}/todo/"
)
