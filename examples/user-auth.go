package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
	"github.com/nelwhix/go-twitter/twitter"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessSecret := os.Getenv("TWITTER_ACCESS_SECRET")

	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	user, _, _ := client.Accounts.VerifyCredentials(verifyParams)
	fmt.Printf("User's ACCOUNT:\n%+v\n", user)

	// Home Timeline
	homeTimelineParams := &twitter.HomeTimelineParams{
		Count:     2,
		TweetMode: "extended",
	}
	tweets, _, _ := client.Timelines.HomeTimeline(homeTimelineParams)
	fmt.Printf("User's HOME TIMELINE:\n%+v\n", tweets)

	// Mention Timeline
	mentionTimelineParams := &twitter.MentionTimelineParams{
		Count:     2,
		TweetMode: "extended",
	}
	tweets, _, _ = client.Timelines.MentionTimeline(mentionTimelineParams)
	fmt.Printf("User's MENTION TIMELINE:\n%+v\n", tweets)

	// Retweets of Me Timeline
	retweetTimelineParams := &twitter.RetweetsOfMeTimelineParams{
		Count:     2,
		TweetMode: "extended",
	}
	tweets, _, _ = client.Timelines.RetweetsOfMeTimeline(retweetTimelineParams)
	fmt.Printf("User's 'RETWEETS OF ME' TIMELINE:\n%+v\n", tweets)

	// Update (POST!) Tweet (uncomment to run)
	tweet, _, _ := client.Statuses.Update("just setting up my twttr", nil)
	fmt.Printf("Posted Tweet\n%v\n", tweet)
}
