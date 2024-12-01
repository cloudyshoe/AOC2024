#!/bin/bash

if [ "$#" -eq 0 ]; then
    echo "Not enough arguments"
    exit 1
fi

if [ "$#" -gt 1 ]; then
    echo "Too many arguments"
    exit 1
fi

if ! [[ "$1" =~ ^[0-9]+$ ]]; then
    echo "Argument must be numeric"
    exit 1
fi

mkdir "$1" &&
cp template/* "$1"/
