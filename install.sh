#!/bin/sh

set -xeuo pipefail

check_go() {
    if ! type go > /dev/null; then
        echo "hello world"
    else
        echo "go exists"
    fi
}

check_go