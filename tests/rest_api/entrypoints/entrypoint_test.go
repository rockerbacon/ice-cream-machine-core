package entrypoints

import (
	assert "rockerbacon/ice-cream-machine-core/pkg/assert"
	entrypoints "rockerbacon/ice-cream-machine-core/internal/rest_api/entrypoints"
	errors "errors"
	http "net/http"
	testing "testing"
)

type MockHttpError struct {}
func (MockHttpError) StatusCode() int {
	return http.StatusBadRequest
}
func (MockHttpError) Error() string {
	return "Testing an error"
}

type MockConnectController struct {}
type MockDeleteController struct {}
type MockGetController struct {}
type MockHeadController struct {}
type MockOptionsController struct {}
type MockPatchController struct {}
type MockPostController struct {}
type MockPutController struct {}
type MockTraceController struct {}
type MockMixedController struct {}
type MockErrorController struct {}

func (MockConnectController) Connect (*http.Request) (any, error) {
	return "tested connect", nil
}
func (MockDeleteController) Delete (*http.Request) (any, error) {
	return "tested delete", nil
}
func (MockGetController) Get (*http.Request) (any, error) {
	return "tested get", nil
}
func (MockHeadController) Head (*http.Request) (any, error) {
	return "tested head", nil
}
func (MockOptionsController) Options (*http.Request) (any, error) {
	return "tested options", nil
}
func (MockPatchController) Patch (*http.Request) (any, error) {
	return "tested patch", nil
}
func (MockPostController) Post (*http.Request) (any, error) {
	return "tested post", nil
}
func (MockPutController) Put (*http.Request) (any, error) {
	return "tested put", nil
}
func (MockTraceController) Trace (*http.Request) (any, error) {
	return "tested trace", nil
}

func (MockMixedController) Delete (*http.Request) (any, error) {
	return "tested mixed delete", nil
}
func (MockMixedController) Get (*http.Request) (any, error) {
	return "tested mixed get", nil
}
func (MockMixedController) Post (*http.Request) (any, error) {
	return "tested mixed post", nil
}

func (MockErrorController) Get (*http.Request) (any, error) {
	return nil, errors.New("Testing an unexpected error")
}
func (MockErrorController) Post (*http.Request) (any, error) {
	return nil, MockHttpError{}
}

type MockResponseWriter struct {
	status *int
	header http.Header
	data *string
}
func makeMockResponseWriter() MockResponseWriter {
	return MockResponseWriter {
		status: new(int),
		header: make(map[string][]string),
		data: new(string),
	}
}
func (self MockResponseWriter) Header() http.Header {
	return self.header
}
func (self MockResponseWriter) WriteHeader(status int) {
	*self.status = status
}
func (self MockResponseWriter) Write(bytes []byte) (int, error) {
	if *self.status == 0 {
		self.WriteHeader(http.StatusOK)
	}

	*self.data += string(bytes)

	return len(bytes), nil
}

func TestBuildsConnectHandler(t *testing.T) {
	writer := makeMockResponseWriter()
	handler := entrypoints.NewHandler(MockConnectController{})
	req := http.Request {
		Method: http.MethodConnect,
	}

	handler.ServeHTTP(writer, &req)

	assert.Equals(t, *writer.data, "\"tested connect\"\n")
	assert.Equals(t, *writer.status, http.StatusOK)
}

func TestBuildsDeleteHandler(t *testing.T) {
	writer := makeMockResponseWriter()
	handler := entrypoints.NewHandler(MockDeleteController{})
	req := http.Request {
		Method: http.MethodDelete,
	}

	handler.ServeHTTP(writer, &req)

	assert.Equals(t, *writer.data, "\"tested delete\"\n")
	assert.Equals(t, *writer.status, http.StatusOK)
}

func TestBuildsGetHandler(t *testing.T) {
	writer := makeMockResponseWriter()
	handler := entrypoints.NewHandler(MockGetController{})
	req := http.Request {
		Method: http.MethodGet,
	}

	handler.ServeHTTP(writer, &req)

	assert.Equals(t, *writer.data, "\"tested get\"\n")
	assert.Equals(t, *writer.status, http.StatusOK)
}

func TestBuildsHeadHandler(t *testing.T) {
	writer := makeMockResponseWriter()
	handler := entrypoints.NewHandler(MockHeadController{})
	req := http.Request {
		Method: http.MethodHead,
	}

	handler.ServeHTTP(writer, &req)

	assert.Equals(t, *writer.data, "\"tested head\"\n")
	assert.Equals(t, *writer.status, http.StatusOK)
}

func TestBuildsOptionsHandler(t *testing.T) {
	writer := makeMockResponseWriter()
	handler := entrypoints.NewHandler(MockOptionsController{})
	req := http.Request {
		Method: http.MethodOptions,
	}

	handler.ServeHTTP(writer, &req)

	assert.Equals(t, *writer.data, "\"tested options\"\n")
	assert.Equals(t, *writer.status, http.StatusOK)
}

func TestBuildsPatchHandler(t *testing.T) {
	writer := makeMockResponseWriter()
	handler := entrypoints.NewHandler(MockPatchController{})
	req := http.Request {
		Method: http.MethodPatch,
	}

	handler.ServeHTTP(writer, &req)

	assert.Equals(t, *writer.data, "\"tested patch\"\n")
	assert.Equals(t, *writer.status, http.StatusOK)
}

func TestBuildsPostHandler(t *testing.T) {
	writer := makeMockResponseWriter()
	handler := entrypoints.NewHandler(MockPostController{})
	req := http.Request {
		Method: http.MethodPost,
	}

	handler.ServeHTTP(writer, &req)

	assert.Equals(t, *writer.data, "\"tested post\"\n")
	assert.Equals(t, *writer.status, http.StatusOK)
}

func TestBuildsPutHandler(t *testing.T) {
	writer := makeMockResponseWriter()
	handler := entrypoints.NewHandler(MockPutController{})
	req := http.Request {
		Method: http.MethodPut,
	}

	handler.ServeHTTP(writer, &req)

	assert.Equals(t, *writer.data, "\"tested put\"\n")
	assert.Equals(t, *writer.status, http.StatusOK)
}

func TestBuildsTraceHandler(t *testing.T) {
	writer := makeMockResponseWriter()
	handler := entrypoints.NewHandler(MockTraceController{})
	req := http.Request {
		Method: http.MethodTrace,
	}

	handler.ServeHTTP(writer, &req)

	assert.Equals(t, *writer.data, "\"tested trace\"\n")
	assert.Equals(t, *writer.status, http.StatusOK)
}

func TestBuildsMixedHandler(t *testing.T) {
	deleteWriter := makeMockResponseWriter()
	getWriter := makeMockResponseWriter()
	postWriter := makeMockResponseWriter()
	patchWriter := makeMockResponseWriter()

	deleteReq := http.Request {
		Method: http.MethodDelete,
	}
	getReq := http.Request {
		Method: http.MethodGet,
	}
	postReq := http.Request {
		Method: http.MethodPost,
	}
	patchReq := http.Request {
		Method: http.MethodPatch,
	}

	handler := entrypoints.NewHandler(MockMixedController{})

	handler.ServeHTTP(deleteWriter, &deleteReq)
	handler.ServeHTTP(getWriter, &getReq)
	handler.ServeHTTP(postWriter, &postReq)
	handler.ServeHTTP(patchWriter, &patchReq)

	assert.Equals(t, *deleteWriter.data, "\"tested mixed delete\"\n")
	assert.Equals(t, *deleteWriter.status, http.StatusOK)
	assert.Equals(t, *getWriter.data, "\"tested mixed get\"\n")
	assert.Equals(t, *getWriter.status, http.StatusOK)
	assert.Equals(t, *postWriter.data, "\"tested mixed post\"\n")
	assert.Equals(t, *postWriter.status, http.StatusOK)
	assert.Equals(t, *patchWriter.data, "{\"error\":\"Method Not Allowed\"}\n")
	assert.Equals(t, *patchWriter.status, http.StatusMethodNotAllowed)
}

func TestHandlesControllerHttpErrors(t *testing.T) {
	writer := makeMockResponseWriter()
	handler := entrypoints.NewHandler(MockErrorController{})
	req := http.Request {
		Method: http.MethodPost,
	}

	handler.ServeHTTP(writer, &req)

	assert.Equals(t, *writer.data, "{\"error\":\"Testing an error\"}\n")
	assert.Equals(t, *writer.status, http.StatusBadRequest)
}

func TestHandlesUnexpectedControllerErrors(t *testing.T) {
	writer := makeMockResponseWriter()
	handler := entrypoints.NewHandler(MockErrorController{})
	req := http.Request {
		Method: http.MethodGet,
	}

	handler.ServeHTTP(writer, &req)

	assert.Equals(t, *writer.data, "{\"error\":\"Testing an unexpected error\"}\n")
	assert.Equals(t, *writer.status, http.StatusInternalServerError)
}

func TestHandlesUnknownHttpMethods(t *testing.T) {
	writer := makeMockResponseWriter()
	handler := entrypoints.NewHandler(MockGetController{})
	req := http.Request {
		Method: "UNKNOWN",
	}

	handler.ServeHTTP(writer, &req)

	assert.Equals(t, *writer.data, "{\"error\":\"Unknown HTTP method\"}\n")
	assert.Equals(t, *writer.status, http.StatusTeapot)
}
