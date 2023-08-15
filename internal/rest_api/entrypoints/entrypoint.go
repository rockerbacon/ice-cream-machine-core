package entrypoints

import (
	api_errors "rockerbacon/ice-cream-machine-core/internal/rest_api/errors"
	errors "errors"
	http "net/http"
	json "encoding/json"
)

type Entrypoint struct {
	Connect func (*http.Request) (any, error)
	Delete func (*http.Request) (any, error)
	Get func (*http.Request) (any, error)
	GetPath func () string
	Head func (*http.Request) (any, error)
	Options func (*http.Request) (any, error)
	Patch func (*http.Request) (any, error)
	Post func (*http.Request) (any, error)
	Put func (*http.Request) (any, error)
	Trace func (*http.Request) (any, error)
}

func NewHandler(e *Entrypoint) http.Handler {
	return http.HandlerFunc(
		func (w http.ResponseWriter, r *http.Request) {
			var responseBody any
			var responseError error

			switch method := r.Method; method {
				case http.MethodConnect:
					responseBody, responseError = e.Connect(r)
				case http.MethodDelete:
					responseBody, responseError = e.Delete(r)
				case http.MethodGet, "":
					responseBody, responseError = e.Get(r)
				case http.MethodHead:
					responseBody, responseError = e.Head(r)
				case http.MethodOptions:
					responseBody, responseError = e.Options(r)
				case http.MethodPatch:
					responseBody, responseError = e.Patch(r)
				case http.MethodPost:
					responseBody, responseError = e.Post(r)
				case http.MethodPut:
					responseBody, responseError = e.Put(r)
				case http.MethodTrace:
					responseBody, responseError = e.Trace(r)
				default:
					responseError = errors.New("Unknown HTTP method")
			}

			httpError, isHttpError := responseError.(api_errors.HttpError)

			var responseStatus int
			if (isHttpError) {
				responseStatus = httpError.StatusCode()
				responseBody = api_errors.ErrorResponseBody{
					Error: httpError.Error(),
				}
			} else if (responseError != nil) {
				responseStatus = http.StatusInternalServerError
				responseBody = api_errors.ErrorResponseBody{
					Error: responseError.Error(),
				}
			} else {
				responseStatus = http.StatusOK
			}

			headers := w.Header()
			headers.Set("content-type", "application/json")
			w.WriteHeader(responseStatus)
			json.NewEncoder(w).Encode(responseBody)
		},
	)
}

