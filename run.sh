#!/usr/bin/env sh
#
go build -o broadband-$(uname)-$(uname -m) *.go
./broadband-$(uname)-$(uname -m)
