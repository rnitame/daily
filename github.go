package main

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"log"

	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	gitconfig "github.com/tcnksm/go-gitconfig"
	"github.com/tidwall/gjson"
	"golang.org/x/oauth2"
)

// NewGitHubClient go-github のクライアント作成
func NewGitHubClient() *github.Client {
	token, err := gitconfig.Global("github.token")
	if err != nil {
		log.Fatalln(errors.Wrap(err, "get token failed"))
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	return github.NewClient(tc)
}

// NewGHEClient go-github のクライアント作成
func NewGHEClient() *github.Client {
	token, err := gitconfig.Global("ghe.token")
	if err != nil {
		log.Fatalln(errors.Wrap(err, "get token failed"))
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	return github.NewClient(tc)
}

// GetEvents GitHub/GHE API から自分のイベントを取得
func GetEvents(client *github.Client, org *string, ghe *string) {
	options := github.ListOptions{Page: 1, PerPage: 50}
	if *ghe != "" {
		baseurl, _ := url.Parse(*ghe)
		client.BaseURL = baseurl
	}
	user, _, err := client.Users.Get(oauth2.NoContext, "")
	if err != nil {
		log.Fatalln(errors.Wrap(err, "get users failed"))
	}

	events, _, err := client.Activity.ListEventsPerformedByUser(oauth2.NoContext, user.GetLogin(), false, &options)
	if err == nil {
		SieveOutEvents(events, org)
	} else {
		log.Fatalln(errors.Wrap(err, "get events failed"))
	}
}

// SieveOutEvents flag によって出すイベントを絞る
func SieveOutEvents(events []*github.Event, org *string) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	today := time.Now()
	const layout = "2006-01-02"
	for _, value := range events {
		// API から取ってきた CreatedAt の文字列に、コマンド叩いた日付が含まれていれば表示
		if strings.Contains(value.CreatedAt.In(jst).String(), string(today.Format(layout))) {
			// organization が指定されていたらその organization のイベントだけ出力
			if *org != "" && !strings.Contains(*value.Repo.Name, *org) {
				continue
			}

			payload, _ := value.ParsePayload()
			// 特定のイベントだけタイトルを表示
			switch *value.Type {
			case "PullRequestEvent":
				pr, ok := payload.(*github.PullRequestEvent)
				if !ok {
					log.Fatalln("Failed type assertion")
				}
				fmt.Println(*value.Repo.Name, *value.Type, *pr.PullRequest.Title)
			case "IssuesEvent":
				issue, ok := payload.(*github.IssuesEvent)
				if !ok {
					log.Fatalln("Failed type assertion")
				}
				fmt.Println(*value.Repo.Name, *value.Type, *issue.Issue.Title)
			default:
				json, _ := value.RawPayload.MarshalJSON()
				action := gjson.Get(string(json), "action")
				fmt.Println(*value.Repo.Name, *value.Type, action)
			}
		}
	}
}
