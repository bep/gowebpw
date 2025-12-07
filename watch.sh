#!/bin/bash

trap exit SIGINT


while true; do ls -d webpwasm/*gz | entr -d touch /Users/bep/dev/go/gohugoio/hugo/resources/image_extended_test.go; done
