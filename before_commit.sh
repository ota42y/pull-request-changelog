#!/bin/bash

echo "go fmt"
find . -path "./src/_vendor" -prune -o -name "*.go" -exec go fmt {} \;

echo "golint"
golint src
