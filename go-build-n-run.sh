#!/bin/bash

# build go-build-n-run application as named "gobr"
go build -o gobr

# move executable file to user command directory
sudo mv gobr /usr/local/bin
