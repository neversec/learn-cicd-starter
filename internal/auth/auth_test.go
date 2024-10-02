package auth

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert" // using testify for assertions
)

func TestGetAPIKey_ValidHeader(t *testing.T) {
	// Define valid headers
	headers := http.Header{
		"Authorization": []string{"ApiKey valid_api_key"},
	}

	// Call the function
	apiKey, err := GetAPIKey(headers)

	// Assert expectations
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, "valid_api_key", apiKey, "api key should be extracted correctly")
}

func TestGetAPIKey_MissingHeader(t *testing.T) {
	// Define empty headers
	headers := http.Header{}

	// Call the function
	apiKey, err := GetAPIKey(headers)

	// Assert expectations
	assert.EqualError(t, err, ErrNoAuthHeaderIncluded.Error(), "expected specific error")
	assert.Empty(t, apiKey, "no api key should be returned")
}

func TestGetAPIKey_MalformedHeader(t *testing.T) {
	// Define malformed headers
	headers := http.Header{
		"Authorization": []string{"Basic invalid_format"},
	}

	// Call the function
	apiKey, err := GetAPIKey(headers)

	// Assert expectations
	assert.EqualError(t, err, "malformed authorization header", "expected specific error")
	assert.Empty(t, apiKey, "no api key should be returned")
}
