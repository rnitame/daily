package main

import (
	"flag"
)

var (
	org = flag.String("org", "", "organization name for showing events")
)

func Run() {
	// ここで flag 受け取って organization の判定
	// 判定によって実行する github.go のメソッドを変える

	flag.Parse()
	client := NewGitHubClient()
	GetEvents(client, org)
}
