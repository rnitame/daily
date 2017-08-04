package main

import (
	"flag"

	"github.com/google/go-github/github"
)

var (
	org = flag.String("org", "", "organization name for showing events")
	ghe = flag.String("ghe", "", "GHE domain")
)

// Run コマンド実行
func Run() {
	// ここで flag 受け取って organization と GHE の判定
	// 判定によって実行する github.go のメソッドを変える

	flag.Parse()
	var client *github.Client

	if *ghe == "" {
		client = NewGitHubClient()
	} else {
		client = NewGHEClient()
	}

	GetEvents(client, org, ghe)
}
