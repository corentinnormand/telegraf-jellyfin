.DEFAULT_GOAL := build

build:
	go build -o bin/jellyfin cmd/main.go

deps:
	go mod vendor

clean:
	rm -r bin

test:
	go test -timeout 30s -count=1 ./plugins/inputs/webhooks/plex

run:
	./bin/jellyfin --config plugin.conf