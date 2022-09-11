#/bin/sh
set -ex

COMMAND="$1"

function linters {
    go vet ./...
}

function tests {
    go test ./...
}

case "$COMMAND" in
    "linters")
        linters;
        ;;
    "tests")
        tests;
        ;;
    "validate")
        linters;
        tests;
        ;;
    "shell")
        sh -i
        ;;
esac
