#!/bin/sh

set -ex

/usr/local/osxcross/bin/lipo \
   -create release/macos/build/cerberusd-arm64 release/macos/build/cerberusd-amd64 \
   -output release/macos/build/cerberusd
