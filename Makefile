.PHONY: build
build:
	go build  -o ./bin/scraper ./cmd/scraper/main.go

.PHONY: build-image
build-image:
	docker build .

.PHONY: run
run: build
	./bin/scraper

.PHONY: test
test:
	go clean -testcache
	go test ./... -v -race -coverprofile=coverage.out -timeout=30m
