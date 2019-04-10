.PHONY: build test

build:
	go build github.com/unjello/commit-activity-log/cmd/commit-activity-log

test:
	go test ./...