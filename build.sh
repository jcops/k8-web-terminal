#!/usr/bin/env bash
GOOS=linux go build -o bin/k8-web-terminal
GOOS=darwin go build -o bin/k8-web-terminal-mac
GOOS=windows go build -o bin/k8-web-terminal.exe