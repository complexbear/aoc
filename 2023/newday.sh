#!/bin/bash

DAY=$1

mkdir "day${DAY}"
cd "day${DAY}"

echo 'def main(testmode bool):\n\tpass' > "main.py"

touch "input.txt"
