#!/bin/bash

version=$(cat VERSION)

go build -ldflags "-X main.BuildVersion=$version"