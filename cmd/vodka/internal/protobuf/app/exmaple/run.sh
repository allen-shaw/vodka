#!/bin/zsh
rm -rf *.go
protoc -I. --go_out=. --go_opt=paths=source_relative --go-gin_out=. --go-gin_opt=paths=source_relative ./hello.proto