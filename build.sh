#!/bin/bash
RUN_NAME=SEproject
mkdir -p output/bin
cp script/* output 2>/dev/null
chmod +x output/bootstrap.sh
go build -ldflags="-s -w" -o "output/bin/${RUN_NAME}"