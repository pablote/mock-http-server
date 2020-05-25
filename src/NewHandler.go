package src

import (
	"log"
	"net/http"
	"net/url"
)

func NewHandler() func(http.ResponseWriter, *http.Request) {
	mockedResponses := map[string]*MockedResponse{}

	return func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("mock-server: log: %v\n", request.RequestURI)
		requestUri, err := url.ParseRequestURI(request.RequestURI)

		if err != nil {
			Send(writer, http.StatusBadRequest, "mock-server: error: Bad request URI")
			return
		}

		mockedResponse, found := mockedResponses[requestUri.Path]

		if IsUpdateCall(*requestUri) {
			if !found {
				mockedResponse = NewDefaultMockedResponse()
				mockedResponses[requestUri.Path] = mockedResponse
			}

			for name, value := range requestUri.Query() {
				if name == "__status" {
					if err := mockedResponse.SetStatus(value); err != nil {
						Send(writer, http.StatusBadRequest, "mock-server: error: invalid port")
					}
				}
			}

			Send(writer, http.StatusOK, "mock-server: route updated")
		} else {
			if !found {
				mockedResponse = NewDefaultMockedResponse()
			}

			Send(writer, mockedResponse.Status.Get(), "")
		}
	}
}
