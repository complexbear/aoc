#!/bin/bash

DAY=$1

mkdir "day${DAY}"
cd "day${DAY}"

cp ../template.py ./main.py
touch "input.txt"
