NAME     := daily
VERSION  := v0.1.0

SRCS    := $(shell find . -type f -name '*.go')

# TODO: set 'make bin/NAME'

test:
	go test -cover -v `glide novendor`

cross-build: 
	for os in darwin linux windows; do \
		for arch in amd64 386; do \
			GOOS=$$os GOARCH=$$arch CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o dist/$$os-$$arch/$(NAME); \
		done; \
	done

create-zip:
	for os in darwin linux windows; do \
		for arch in amd64 386; do \
			zip -r -m -q dist/$(NAME)-$$os-$$arch.zip dist/$$os-$$arch; \
		done; \
	done
