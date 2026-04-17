set dotenv-load

export PROJECTDIR := justfile_directory()
export BINARYCG := justfile_directory() + "/cmd/binary"

default:
    @just --list

generate:
    @go generate ./...

build: generate
    @mkdir -p ./build
    @go build -o ./build/gs ./cmd/gs

build-dist: generate
    @mkdir -p ./dist
    GOOS=linux GOARCH=amd64 go build -o ./dist/gs-linux-amd64 ./cmd/gs
    GOOS=darwin GOARCH=amd64 go build -o ./dist/gs-darwin-amd64 ./cmd/gs
    GOOS=darwin GOARCH=arm64 go build -o ./dist/gs-darwin-arm64 ./cmd/gs
    GOOS=windows GOARCH=amd64 go build -o ./dist/gs-windows-amd64.exe ./cmd/gs

install: build
    @cp ./build/gs $HOME/bin/

test: generate
    @go test -race ./tests/unit/parser
    @go test -race ./tests/unit/value
    @go test -race ./tests/unit/stdlib/json
    @go test -race ./tests/unit/stdlib
    @go test -race ./tests/unit
    @go run ./cmd/gs -resolve ./tests/testdata/cli/test.gs

bench-tool: generate
    @go run ./cmd/bench

clean:
    rm -rf ./build
    rm -rf ./dist
    rm -rf ./*.prof
    rm -rf ./*.log

bench-test: generate
    @go test -test.fullpath=true -run ^$ -bench=^BenchmarkVM$ -benchmem -cpuprofile cpu.prof -memprofile mem.prof -trace trace.prof ./tests/benchmark

cpu:
       go tool pprof -http=:8080 cpu.prof

mem:
    go tool pprof -http=:8080 mem.prof

trace:
       go tool trace trace.prof