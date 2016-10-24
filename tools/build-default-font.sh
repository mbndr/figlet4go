#!/usr/bin/env bash
# Build the bindata.go file with the default font in it
cd ../assets
go-bindata -o ../bindata.go -pkg figlet4go ./