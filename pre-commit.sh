#!/bin/sh -eux

goimports -w .
go tool vet .

if ! golint ./... | grep -v "\-gen.go"; then
    echo "golint ok!"
else
    exit 1
fi
