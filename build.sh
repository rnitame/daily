#!/bin/sh

if [ $# != 1 ]; then
	echo "Usage: $0 [distary name]"
	exit 0
fi

GOOS=linux GOARCH=amd64 go build -o ./dist/$1_linux_amd64/$1
GOOS=linux GOARCH=386 go build -o ./dist/$1_linux_386/$1

GOOS=windows GOARCH=386 go build -o ./dist/$1_windows_386/$1.exe
GOOS=windows GOARCH=amd64 go build -o ./dist/$1_windows_amd64/$1.exe

GOOS=darwin GOARCH=386 go build -o ./dist/$1_darwin_386/$1
GOOS=darwin GOARCH=amd64 go build -o ./dist/$1_darwin_amd64/$1