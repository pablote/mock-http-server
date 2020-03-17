package src

import (
	"net/url"
	"strings"
)

func IsUpdateCall(requestUri url.URL) bool {
	for name, _ := range requestUri.Query() {
		if strings.HasPrefix(name, "__") {
			return true
		}
	}

	return false
}
