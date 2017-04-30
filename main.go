package main

import (
	"fmt"
	"log"
	"os"
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

		fmt.Print(string(buf[:n]))
	}
}
