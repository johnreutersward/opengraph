#!/bin/bash

APP=opengraph
ARCH=amd64

# Set version
read -p "Version: " VERSION

# Build for macOS
GOOS=darwin GOARCH=$ARCH go build
zip $APP-$VERSION-darwin-$ARCH.zip $APP

# Build for Linux
GOOS=linux GOARCH=$ARCH go build
zip $APP-$VERSION-linux-$ARCH.zip $APP

# Build for Windows
GOOS=windows GOARCH=$ARCH go build
zip $APP-$VERSION-windows-$ARCH.zip $APP.exe

# Clean up
rm $APP
rm $APP.exe
