package handler

const (
	ErrRequestInvalid    = 40001
	ErrInvalidPassword   = 40101
	ErrTokenRequired     = 40102
	ErrInvalidToken      = 40103
	ErrUnauthorizedToken = 40104
	ErrForbidden         = 40301
	ErrNotFound          = 40401
	ErrInternalServer    = 50001

	ErrInternalServerMsg = "internal server error"
)

type Response struct {
	Data any `json:"data"`
}

type ErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
