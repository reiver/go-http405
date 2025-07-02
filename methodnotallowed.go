package http405

import (
	"net/http"
)

func MethodNotAllowed(responseWriter http.ResponseWriter, request *http.Request, methods ...string) {
	Serve(responseWriter, methods...)
}
