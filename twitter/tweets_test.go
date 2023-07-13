package twitter

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTweetService_PostTweetNilParams(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()
	mux.HandleFunc("/2/tweets.json", func(w http.ResponseWriter, r *http.Request) {
		assertPostForm(t, map[string]string{"text": "very informative tweet"}, r)
	})
	client := NewClient(httpClient)
	_, _, err := client.Tweets.PostTweet("very informative tweet", nil)
	
	assert.Nil(t, err)
}

// func TestStatusService_HTTPError(t *testing.T) {
// 	httpClient, _, server := testServer()
// 	server.Close()
// 	client := NewClient(httpClient)
// 	_, _, err := client.Tweets.PostTweet("very informative tweet", nil)
// 	if err == nil || !strings.Contains(err.Error(), "connection refused") || !strings.Contains(err.Error(), "No connection could be made") {
// 		t.Errorf("Statuses.Update error expected connection refused, got: \n %+v", err)
// 	}
// }
