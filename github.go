package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/go-github/github"
	gitconfig "github.com/tcnksm/go-gitconfig"
	"github.com/tidwall/gjson"
	"golang.org/x/oauth2"
)

func NewGitHubClient() *github.Client {
	// グローバルな gitconfig にあるトークンを持ってくる
	token, err := gitconfig.Global("github.token")
	if err != nil {
		log.Fatal(err)
	}

	// go-github と oauth2 で GitHub の認証
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	return github.NewClient(tc)
}

func GetEvents(client *github.Client, org *string) {
	options := github.ListOptions{Page: 1, PerPage: 50}
	user, _, err := client.Users.Get(oauth2.NoContext, "")
	events, _, err := client.Activity.ListEventsPerformedByUser(oauth2.NoContext, user.GetLogin(), false, &options)
	if err == nil {
		SieveOutEvents(events, org)
	} else {
		log.Fatal(err)
	}
}

// イベントのふるい分け
func SieveOutEvents(events []*github.Event, org *string) {
	// コマンド叩いた日のイベントを表示する
	jst, _ := time.LoadLocation("Asia/Tokyo")
	today := time.Now()
	const layout = "2006-01-02"
	for _, value := range events {
		// API から取ってきた CreatedAt の文字列に、コマンド叩いた日付が含まれていれば表示
		if strings.Contains(value.CreatedAt.In(jst).String(), string(today.Format(layout))) {
			json, _ := value.RawPayload.MarshalJSON()
			payload := gjson.Get(string(json), "action")

			// organization が指定されていたらその organization のイベントだけ出力
			if *org != "" && !strings.Contains(*value.Repo.Name, *org) {
				continue
			}
			fmt.Println(*value.Repo.Name, *value.Type, payload)
		}
	}

}
