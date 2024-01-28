package entrypoints

import (
	api_errors "rockerbacon/ice-cream-machine-core/internal/rest_api/errors"
	common "rockerbacon/ice-cream-machine-core/internal/rest_api/common_controllers"
	http "net/http"
	json "encoding/json"
)

type ConnectController interface {
	Connect (*http.Request) (any, error)
}
type DeleteController interface {
	Delete (*http.Request) (any, error)
}
type GetController interface {
	Get (*http.Request) (any, error)
}
type HeadController interface {
	Head (*http.Request) (any, error)
}
type OptionsController interface {
	Options (*http.Request) (any, error)
}
type PatchController interface {
	Patch (*http.Request) (any, error)
}
type PostController interface {
	Post (*http.Request) (any, error)
}
type PutController interface {
	Put (*http.Request) (any, error)
}
type TraceController interface {
	Trace (*http.Request) (any, error)
}

type EntrypointController struct {
	Connector ConnectController
	Deletor DeleteController
	Getter GetController
	Header HeadController
	Optioner OptionsController
	Patcher PatchController
	Poster PostController
	Putter PutController
	Tracer TraceController
}

func convertController[Controller any](controller any) Controller {
	castController, controllerImplementsMethod := controller.(Controller)
	if controllerImplementsMethod {
		return castController
	}

	return (any)(common.MethodNotAllowed).(Controller)
}

func NewHandler(controller any) http.Handler {
	e := EntrypointController {
		Connector: convertController[ConnectController](controller),
		Deletor: convertController[DeleteController](controller),
		Getter: convertController[GetController](controller),
		Header: convertController[HeadController](controller),
		Optioner: convertController[OptionsController](controller),
		Patcher: convertController[PatchController](controller),
		Poster: convertController[PostController](controller),
		Putter: convertController[PutController](controller),
		Tracer: convertController[TraceController](controller),
	}

	return http.HandlerFunc(
		func (w http.ResponseWriter, r *http.Request) {
			var responseBody any
			var responseError error

			switch method := r.Method; method {
				case http.MethodConnect:
					responseBody, responseError = e.Connector.Connect(r)
				case http.MethodDelete:
					responseBody, responseError = e.Deletor.Delete(r)
				case http.MethodGet, "":
					responseBody, responseError = e.Getter.Get(r)
				case http.MethodHead:
					responseBody, responseError = e.Header.Head(r)
				case http.MethodOptions:
					responseBody, responseError = e.Optioner.Options(r)
				case http.MethodPatch:
					responseBody, responseError = e.Patcher.Patch(r)
				case http.MethodPost:
					responseBody, responseError = e.Poster.Post(r)
				case http.MethodPut:
					responseBody, responseError = e.Putter.Put(r)
				case http.MethodTrace:
					responseBody, responseError = e.Tracer.Trace(r)
				default:
					responseError = api_errors.UnknownHttpMethodError{}
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

