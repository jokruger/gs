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

test: generate
    @go test -race ./tests/unit/parser
    @go test -race ./tests/unit/value
    @go test -race ./tests/unit/stdlib/json
    @go test -race ./tests/unit/stdlib
    @go test -race ./tests/unit
    @go run ./cmd/gs -resolve ./tests/testdata/cli/test.gs

clean:
    rm -rf ./build
    rm -rf ./*.prof
    rm -rf ./*.log
