#!/bin/bash
if [ -z "$1" ]; then
    echo "Usage: ./release.sh <tag_name>"
    exit 1
fi

git tag "$1"
git push origin "$1"
GOPROXY=proxy.golang.org go list -m github.com/Fesaa/enka-network-api-go@"$1"
