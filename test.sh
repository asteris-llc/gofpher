#!/usr/bin/env bash

set -e

for d in $(go list ./...); do
    go test -race "$d"
done
