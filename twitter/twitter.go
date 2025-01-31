package twitter

import (
	"net/http"

	"github.com/dghubble/sling"
)

const API_V1 = "https://api.twitter.com/1.1/"
const API_V2 = "https://api.twitter.com/2/"

// Client is a Twitter client for making Twitter API requests.
type Client struct {
	sling *sling.Sling
	// Twitter API Services
	Tweets      *TweetService
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base(API_V2)
	
	return &Client{
		sling:          base,
		Tweets:       newTweetService(base.New()),
	}
}

// Bool returns a new pointer to the given bool value.
func Bool(v bool) *bool {
	ptr := new(bool)
	*ptr = v
	return ptr
}

// Float returns a new pointer to the given float64 value.
func Float(v float64) *float64 {
	ptr := new(float64)
	*ptr = v
	return ptr
}
