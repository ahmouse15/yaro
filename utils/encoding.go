package utils

import (
	"encoding/base64"
	"strings"
)

// Encodes given string to url-safe Base64
func ToBase64Url(input string) string {
	buffer := new(strings.Builder)

	encoder := base64.NewEncoder(base64.URLEncoding, buffer)

	encoder.Write([]byte(input))
	encoder.Close()

	// Remove padding (=)
	output := strings.Replace(buffer.String(), "=", "", 2)

	return output
}

// Probably not needed
func fromBase64Url() {

}
