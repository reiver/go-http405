package http405

import (
	"net/http"
)

func Serve(responseWriter http.ResponseWriter, methods ...string) error {
	return ServeString(responseWriter, DefaultStatusText, methods...)
}
