package src

import (
	"fmt"
	"net/http"
)

func Send(writer http.ResponseWriter, status int, format string, a ...interface{}) {
	writer.WriteHeader(status)
	_, _ = fmt.Fprintf(writer, format, a...)
}
