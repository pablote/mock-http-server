package main

import (
	"fmt"
	"github.com/pablote/mock-server/src"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func Send(writer http.ResponseWriter, status int, format string, a ...interface{}) {
	writer.WriteHeader(status)
	_, _ = fmt.Fprintf(writer, format, a...)
}

func main() {
	port := "9090"
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}

	mockedResponses := map[string]*src.MockedResponse{}

	handler := func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("mock-server: log: %v\n", request.RequestURI)
		requestUri, err := url.ParseRequestURI(request.RequestURI)

		if err != nil {
			Send(writer, http.StatusBadRequest, "mock-server: error: Bad request URI")
			return
		}

		mockedResponse, found := mockedResponses[requestUri.Path]

		if src.IsUpdateCall(*requestUri) {
			if !found {
				mockedResponse = src.NewDefaultMockedResponse()
				mockedResponses[requestUri.Path] = mockedResponse
			}

			for name, value := range requestUri.Query() {
				if name == "__status" {
					newStatus, err := strconv.ParseInt(value[0], 0, 32)
					if err != nil {
						Send(writer, http.StatusBadRequest, "mock-server: error: invalid port")
						return
					}

					mockedResponse.SetStatus(int(newStatus))
				}
			}

			Send(writer, http.StatusOK, "mock-server: config updated")
		} else {
			if found {
				Send(writer, mockedResponse.Status, "")
			} else {
				Send(writer, http.StatusOK, "")
			}
		}
	}

	log.Printf("mock-server starting on port %v\n", port)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}