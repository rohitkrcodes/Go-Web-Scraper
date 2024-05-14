package auth

import (
	"errors"
	"net/http"
	"strings"
)

// extracts an APIKey from
// headers of an http request
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 && vals[0] != "ApiKey" {
		return "", errors.New("invalid auth header") // Format: Authorization: APIKey {insert apikey here}
	}

	return vals[1], nil
}
