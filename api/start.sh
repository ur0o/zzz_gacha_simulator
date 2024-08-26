#! /bin/sh

go mod tidy

npm run dev &
air -c air.toml
