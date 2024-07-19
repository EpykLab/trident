#!/bin/bash

# Update the package list
sudo apt update

# Install the Protocol Buffers Compiler
sudo apt install -y protobuf-compiler

# Verify the installation of protoc
protoc --version

# Install the Go plugins for protoc
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Update the PATH to include the Go bin directory (if not already included)
export PATH="$PATH:$(go env GOPATH)/bin"

# Print the installation paths of the Go plugins
which protoc-gen-go
which protoc-gen-go-grpc

echo "Protocol Buffers setup is complete."
