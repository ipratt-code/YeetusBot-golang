package reddit

import (
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"fmt"
	"math/rand"
	"context"
	"strings"
	//"os"
)

func redditRandomRetrieve(subreddits []string, randomRange int) (*reddit.Post, error){
    stringedSubreddits := strings.Join(subreddits, "+")
	//credentials := reddit.Credentials{ID: os.Getenv("REDDIT_BOT_ID"), Secret: os.Getenv("REDDIT_BOT_SECRET"), Username: "username", Password: "password"}
    client, _ := reddit.NewReadonlyClient() //reddit.NewClient(credentials)
	posts, _, err := client.Subreddit.TopPosts(context.Background(), stringedSubreddits, &reddit.ListPostOptions{
		ListOptions: reddit.ListOptions{
			Limit: randomRange,
		},
		Time: "all",
	})
	if err != nil {
		return nil, err
	}else if len(posts) < 1 {
		return nil, fmt.Errorf("There are no posts here!")
	}
	p := posts[rand.Intn(len(posts))]
	if p.NSFW == true {
		return nil, fmt.Errorf("This is a NSFW post!")
	}
	return p, err
}