test:
	go test ./...

bench:
	go test -bench=. ./...

run:
	go run .

build:
	go build -o bin/aoc2018 .