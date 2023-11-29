#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=SEproject
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}
