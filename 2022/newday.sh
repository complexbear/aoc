#!/bin/bash

DAY=$1

mkdir "day${DAY}"
cd "day${DAY}"

echo "package day${DAY}\n\nfunc Main(testmode bool) {}" > "day${DAY}.go"

touch "day${DAY}.txt"
