#!/bin/bash

echo "Fetching version"
version=$(cat VERSION)

echo "Building $version of ChiaBoard"
go build -ldflags "-X main.BuildVersion=$version"

echo "Build finished"