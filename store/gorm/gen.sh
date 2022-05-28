#!/bin/bash
set -e
set -x

go mod tidy
go build -o store/gorm/gorm-gen ./store/gorm/cmd/gen/main.go

cd store/gorm/migrations
go install github.com/pressly/goose/v3/cmd/goose@latest
goose mysql "$MYSQL" up

cd ..
./gorm-gen -mysql "$MYSQL" -db "$DB"
