package twitter

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

// Tweet represents a Twitter Tweet, previously called a status.
// I added the fields in use now, more fields will be added as functionality increases
// https://developer.twitter.com/en/docs/twitter-api/data-dictionary/object-model/tweet
type Tweet struct {
	ID                   string                  `json:"id"`
	Text                 string                 `json:"text"`
	AuthorID			 string					`json:"author_id"`
	ConversationID 		 string                 `json:"conversation_id"`
	CreatedAt            string                 `json:"created_at"`	
	InReplyToUserID      string                  `json:"in_reply_to_user_id"`
}

// CreatedAtTime returns the time a tweet was created.
func (t Tweet) CreatedAtTime() (time.Time, error) {
	return time.Parse(time.RubyDate, t.CreatedAt)
}

// TweetService provides methods for accessing Twitter status API endpoints.
type TweetService struct {
	sling *sling.Sling
}

// newTweetService returns a new TweetService.
func newTweetService(sling *sling.Sling) *TweetService {
	return &TweetService{
		sling: sling.Path(""),
	}
}

// StatusUpdateParams are the parameters for StatusService.Update
type PostTweetParams struct {
	DirectMessageDeepLink	string 			`json:"direct_message_deep_link,omitempty"`
	ForSuperFollowersOnly 	bool 			`json:"for_super_followers_only,omitempty"`
	QuoteTweetId			string 			`json:"quote_tweet_id,omitempty"`
	Text 					string 			`json:"text,omitempty"`
	Media					*TweetMedia		`json:"media,omitempty"`
}

type TweetMedia struct {
	MediaIds 			[]string		`json:"media_ids,omitempty"`
	TaggedUserIds		[]string		`json:"tagged_user_ids,omitempty"`
}


// Post tweet makes a new tweet.
// Requires a user auth context.
// https://developer.twitter.com/en/docs/twitter-api/v1/tweets/post-and-engage/api-reference/post-statuses-update
func (s *TweetService) PostTweet(text string, params *PostTweetParams) (*Tweet, *http.Response, error) {
	if params == nil {
		params = &PostTweetParams{}
	}
	params.Text = text
	tweet := new(Tweet)
	apiError := new(APIError)

	resp, err := s.sling.New().Post("tweets").BodyJSON(params).Receive(tweet, apiError)
	return tweet, resp, relevantError(err, *apiError)
}

// Destroy deletes the Tweet with the given id and returns it if successful.
// Requires a user auth context.
// https://developer.twitter.com/en/docs/twitter-api/tweets/manage-tweets/api-reference/delete-tweets-id
func (s *TweetService) Destroy(id string) (*Tweet, *http.Response, error) {
	tweet := new(Tweet)
	apiError := new(APIError)
	path := fmt.Sprintf("/tweets/%v", id)
	resp, err := s.sling.New().Delete(path).Receive(tweet, apiError)

	return tweet, resp, relevantError(err, *apiError)
}

// GetTweetsParams are the parameters for TweetService.Lookup
type GetTweetsParams struct {
	Ids		[]string 	`url:"ids,omitempty,comma"`
}

// Lookup returns the requested Tweets as a slice. Combines ids from the
// required ids argument and from params.Id.
// https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets
func (s *TweetService) GetTweets(ids []string, params *GetTweetsParams) ([]Tweet, *http.Response, error) {
	if params == nil {
		params = &GetTweetsParams{}
	}
	params.Ids = append(params.Ids, ids...)
	tweets := new([]Tweet)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/tweets").QueryStruct(params).Receive(tweets, apiError)

	return *tweets, resp, relevantError(err, *apiError)
}

// Show returns the requested Tweet.
// https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets-id
func (s *TweetService) Show(id string) (*Tweet, *http.Response, error) {
	tweet := new(Tweet)
	apiError := new(APIError)
	path := fmt.Sprintf("/tweets/%v", id)
	resp, err := s.sling.New().Get(path).Receive(tweet, apiError)

	return tweet, resp, relevantError(err, *apiError)
}

// StatusRetweeterParams are the parameters for StatusService.Retweeters.
type TweetRetweeterParams struct {
	MaxResults     int64 `url:"max_results,omitempty"`
}

// RetweeterIDs is a cursored collection of User IDs that retweeted a Tweet.
type Retweeter struct {
	ID					string 		`json:"id"`
	Name				string		`json:"name"`
	Username			string		`json:"username"`
	CreatedAt			string 		`json:"created_at,omitempty"`
	Description			string 		`json:"description,omitempty"`
	PinnedTweetId		string 		`json:"pinned_tweet_id,omitempty"`
}

// Retweeters return the retweeters of a specific tweet.
// https://developer.twitter.com/en/docs/tweets/post-and-engage/api-reference/get-statuses-retweeters-ids
func (s *TweetService) Retweeters(id string, params *TweetRetweeterParams) (*[]Retweeter, *http.Response, error) {
	retweeters := new([]Retweeter)
	apiError := new(APIError)
	path := fmt.Sprintf("/tweets/%v/retweeted_by", id)
	resp, err := s.sling.Get(path).QueryStruct(params).Receive(retweeters, apiError)

	return retweeters, resp, relevantError(err, *apiError)
}

// StatusRetweetsParams are the parameters for StatusService.Retweets
type RetweetParams struct {
	TweetID 		string		`json:"tweet_id"`
}

// Retweets returns the most recent retweets of the Tweet with the given id.
// https://developer.twitter.com/en/docs/twitter-api/tweets/retweets/api-reference/post-users-id-retweets
func (s *TweetService) Retweet(userId string, params *RetweetParams) (*Tweet, *http.Response, error) {
	tweet := new(Tweet)
	apiError := new(APIError)
	path := fmt.Sprintf("/users/%v/retweets", userId)
	resp, err := s.sling.New().Post(path).BodyJSON(params).Receive(tweet, apiError)

	return tweet, resp, relevantError(err, *apiError)
}



