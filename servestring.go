package http405

import (
	"net/http"
	"io"
)

func ServeString(responseWriter http.ResponseWriter, value string, methods ...string) error {
	if nil == responseWriter {
		return ErrNilHTTPResponseWriter
	}

	responseWriter.Header().Add("Allow", allow(methods...))

	responseWriter.WriteHeader(StatusCode)

	if "" != value {
		io.WriteString(responseWriter, value)
	}

	return nil
}
