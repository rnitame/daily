NAME     := daily
VERSION  := v0.3.0

SRCS    := $(shell find . -type f -name '*.go')

# todo: make bin/NAME

.PHONY: test
	test:
	go test -cover -v `glide novendor`

.PHONY: cross-build
	cross-build: deps
	for os in darwin linux windows; do \
		for arch in amd64 386; do \
		GOOS=$$os GOARCH=$$arch CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o dist/$$os-$$arch/$(NAME); \
		done; \
		done

DIST_DIRS := find * -type d -exec

.PHONY: dist
	dist:
		cd dist && \
		$(DIST_DIRS) cp ../LICENSE {} \; && \
		$(DIST_DIRS) cp ../README.md {} \; && \
		$(DIST_DIRS) tar -zcf $(NAME)-$(VERSION)-{}.tar.gz {} \; && \
		$(DIST_DIRS) zip -r $(NAME)-$(VERSION)-{}.zip {} \; && \
		cd ..
