package api_errors

type HttpError interface {
	StatusCode() int
	Error() string
}

type ErrorResponseBody struct {
	Error string `json:"error"`
}

