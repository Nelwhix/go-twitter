package twitter

import (
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
	Geo						TweetGeoObj 	`json:"geo,omitempty"`
	Media 					TweetMedia 		`json:"media,omitempty"`
	Poll					TweetPoll		`json:"poll,omitempty"`
	QuoteTweetId			string 			`json:"quote_tweet_id,omitempty"`
	Reply					TweetReply		`json:"reply,omitempty"`
	ReplySettings			ReplySetting	`json:"reply_settings,omitempty"`
	Text 					string 			`json:"text,omitempty"`
}

type ReplySetting string

const (
    MentionedUsers ReplySetting = "mentionedUsers"
    Following      ReplySetting = "following"
)

type TweetGeoObj struct {
	PlaceId		string `json:"place_id,omitempty"`
}

type TweetMedia struct {
	MediaIds 			[]string		`json:"media_ids,omitempty"`
	TaggedUserIds		[]string		`json:"tagged_user_ids,omitempty"`
}

type TweetPoll struct {
	DurationMinutes		int			`json:"duration_minutes,omitempty"`
	Options				[]string	`json:"options,omitempty"`

}

type TweetReply struct {
	ExcludeReplyUserIds			[]string		`json:"exclude_reply_user_ids,omitempty"`
	InReplyToTweetId			string 			`json:"in_reply_to_tweet_id,omitempty"`
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


