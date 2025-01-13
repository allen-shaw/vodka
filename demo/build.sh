#!/bin/bash
SERVER_NAME=demo_server

rm -rf out
mkdir out
mkdir out/bin
mkdir out/conf

go build -o $SERVER_NAME ./main.go
chmod +x $SERVER_NAME
mv $SERVER_NAME out/bin
cp -r conf out/conf
cp scripts/bootstrap.sh out
