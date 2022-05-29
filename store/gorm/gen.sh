#!/bin/bash
# the script should be run from the project root
# use goose to migrate the database and use gorm/gen to generate the models in go code
set -e
set -x

go mod tidy
go build -o store/gorm/gorm-gen ./store/gorm/cmd/gen/main.go

cd store/gorm/migrations
go install github.com/pressly/goose/v3/cmd/goose@latest
goose mysql "$MYSQL" up

cd ..
./gorm-gen -conn "$MYSQL" -db "$DB"
