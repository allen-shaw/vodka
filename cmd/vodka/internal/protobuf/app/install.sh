#!/bin/zsh
EXE=protoc-gen-go-gin
go build -tags debug -o $EXE main.go 
chmod +x $EXE
sudo mv $EXE $GOPATH/bin/