package main

import (
	"log"
	"os"

	"fmt"

	"bufio"

	"time"

	"strings"

	"github.com/google/go-github/github"
	"github.com/tidwall/gjson"
	"golang.org/x/oauth2"
)

// 読み込みバッファのサイズ
const (
	BUFSIZE = 1024
)

func main() {
	// ファイルを読み込んでトークン取得
	file, err := os.Open(`./token.txt`)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := bufio.NewScanner(file)
	token := ""
	for buf.Scan() {
		token = buf.Text()
	}
	if err := buf.Err(); err != nil {
		log.Fatal(err)
	}

	// go-github と oauth2 で GitHub の認証
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)
	options := github.ListOptions{Page: 1, PerPage: 50}
	user, _, err := client.Users.Get(oauth2.NoContext, "")
	events, _, err := client.Activity.ListEventsPerformedByUser(oauth2.NoContext, user.GetLogin(), false, &options)
	if err != nil {
		log.Fatal(err)
	}

	// コマンド叩いた日のイベントを表示する
	jst, _ := time.LoadLocation("Asia/Tokyo")
	today := time.Now()
	const layout = "2006-01-02"
	for _, value := range events {
		// API から取ってきた CreatedAt の文字列に、コマンド叩いた日付が含まれていれば表示
		if strings.Contains(value.CreatedAt.In(jst).String(), string(today.Format(layout))) {
			json, _ := value.RawPayload.MarshalJSON()
			payload := gjson.Get(string(json), "action")
			fmt.Println(*value.Repo.Name, *value.Type, payload)
		}
	}
}
