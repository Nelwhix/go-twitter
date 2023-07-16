package twitter

import (
	"fmt"
)

// APIError represents a Twitter API Error response
type APIError struct {
	Title		string		`json:"title"`
	Type 		string 		`json:"type"`
	Status 		int			`json:"status"`
	Detail		string 		`json:"detail"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("Error sending request: %d %v", e.Status, e.Detail)
}

// relevantError returns any non-nil http-related error (creating the request,
// getting the response, decoding) if any. If the decoded apiError is non-zero
// the apiError is returned. Otherwise, no errors occurred, returns nil.
func relevantError(httpError error, apiError APIError) error {
	if httpError != nil {
		return httpError
	}

	if apiError.Detail == "" {
		return nil
	}
	return apiError
}