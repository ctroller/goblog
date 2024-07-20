#!/usr/bin/env bash

tailwindcss -i ui/src/css/main.css -o ui/static/css/tailwind.css --minify
go run cmd/blog/server.go