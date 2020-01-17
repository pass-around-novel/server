#!/bin/bash
set -e
cd "$(dirname "$0")"

sed -f watch-gen.sed watch-gen.csv > watch-gen.go
