package main

import (
	"log"
	"os"

	"fmt"

	"bufio"

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

	buf := bufio.NewScanner(file)
	token := ""
	for buf.Scan() {
		token = buf.Text()
	}
	if err := buf.Err(); err != nil {
		log.Fatal(err)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)
	repos, _, err := client.Repositories.List(oauth2.NoContext, "", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(repos)
}
