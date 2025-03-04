.PHONY: build
build:
	go build  -o ./bin/scraper ./cmd/scraper/main.go

.PHONY: build-image
build-image:
	docker build .

.PHONY: run
run: build
	./bin/scraper
