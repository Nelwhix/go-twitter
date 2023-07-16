package twitter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var errAPI = APIError{
	Type: "about:blank",
	Title: "Unauthorized",
	Status: 401,
	Detail: "Unauthorized",
}
var errHTTP = fmt.Errorf("unknown host")

func TestAPIError_Error(t *testing.T) {
	err := APIError{}
	if assert.Error(t, err) {
		assert.Equal(t, "Error sending request: 0 ", err.Error())
	}
	if assert.Error(t, errAPI) {
		assert.Equal(t, "Error sending request: 401 Unauthorized", errAPI.Error())
	}
}


func TestRelevantError(t *testing.T) {
	cases := []struct {
		httpError error
		apiError  APIError
		expected  error
	}{
		{nil, APIError{}, nil},
		{nil, errAPI, errAPI},
		{errHTTP, APIError{}, errHTTP},
		{errHTTP, errAPI, errHTTP},
	}
	for _, c := range cases {
		err := relevantError(c.httpError, c.apiError)
		assert.Equal(t, c.expected, err)
	}
}