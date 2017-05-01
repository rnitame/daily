package main

import (
	"context"
	"log"
	"os"

	"fmt"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// 読み込みバッファのサイズ
const (
	BUFSIZE = 1024
)

func main() {
	file, err := os.Open(`./token.txt`)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, BUFSIZE)
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			log.Fatal(err)
			break
		}

	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ""},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	repos, _, err := client.Repositories.List(ctx, "", nil)

	fmt.Print(repos)
}
