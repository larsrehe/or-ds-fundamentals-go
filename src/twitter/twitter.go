package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

var (
	consumerKey    = os.Getenv("CONSUMER_KEY_TWITTER")
	consumerSecret = os.Getenv("CONSUMER_SECRET_TWITTER")
	accessToken    = os.Getenv("ACCESS_KEY_TWITTER")
	accessSecret   = os.Getenv("ACCESS_SECRET_TWITTER")
)

type cleanTweet struct {
	Id       string
	Text     string
	Likes    int
	Retweets int
	Language string
	URL      string
}

var CleanTweets []cleanTweet

func getCleanTweets(tweet anaconda.Tweet) cleanTweet {
	var t = cleanTweet{tweet.IdStr, tweet.Text,
		tweet.FavoriteCount, tweet.RetweetCount,
		tweet.Lang, "www.twitter.com/i/web/status/" + tweet.IdStr}
	return t
}

func PrettyPrintTweet(tweet anaconda.Tweet) {
	t := getCleanTweets(tweet)
	tweetJSON, _ := json.MarshalIndent(t, "", "\t")
	fmt.Println(string(tweetJSON))
}

func main() {
	api := anaconda.NewTwitterApiWithCredentials(
		accessToken, accessSecret, consumerKey, consumerSecret)
	fmt.Println("Started the API...")

	searchResult, _ := api.GetSearch("deep learning",
		url.Values{"result_type": []string{"popular"}})

	fmt.Printf("Retrieved %v tweets\n",
		len(searchResult.Statuses))

	for _, tweet := range searchResult.Statuses {
		if tweet.FavoriteCount > 2000 && tweet.RetweetCount > 500 {
			_, err := api.Retweet(tweet.Id, false)
			if err != nil {
				fmt.Println("Error in Retweeting")
				continue
			}
		} else {
			PrettyPrintTweet(tweet)
		}
	}
}
