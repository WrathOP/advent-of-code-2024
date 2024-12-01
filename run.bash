#!/bin/bash

# Check if day parameter is provided
if [ $# -eq 0 ]; then
    echo "Usage: ./run.sh <day> <part>"
    echo "Example: ./run.sh 1 1"
    exit 1
fi

go build && ./advent-of-code-2024 -day "$1" -part "$2"
