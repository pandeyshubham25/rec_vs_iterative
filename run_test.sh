#!/bin/bash

# Read command line arguments
while [[ $# -gt 0 ]]; do
    key="$1"
    case $key in
        -depth)
            DEPTH="$2"
            shift 2
            ;;
        -maxVariables)
            MAXVARIABLES="$2"
            shift 2
            ;;
        -step)
            STEP="$2"
            shift 2
            ;;
        -fileName)
            FILENAME="$2"
            shift 2
            ;;
        *)
            echo "Unknown option: $1"
            exit 1
            ;;
    esac
done

# Run the Go program with the provided arguments
go run rec.go -depth "$DEPTH" -maxVariables "$MAXVARIABLES" -step "$STEP" -fileName "$FILENAME"
python plot.py "$FILENAME"