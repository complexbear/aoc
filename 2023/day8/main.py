import operator
from io import StringIO
from itertools import cycle
from math import lcm
from functools import reduce
from typing import Dict

EXAMPLE_1 = """LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)"""

EXAMPLE_2 = """LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)"""

Directions = str
Path = Dict[str, tuple]


def parse(text) -> (Directions, Path):
    directions = text.readline().strip()
    text.readline()

    path = {}
    for line in text.readlines():
        key, steps = line.split(" = ")
        steps = (steps[1:4], steps[6:9])
        path[key] = steps

    return directions, path


def part1(text) -> int:
    directions, path = parse(text)
    directions = cycle(directions)
    i = 0
    key = "AAA"
    while key != "ZZZ":
        d = next(directions)
        key = path[key][0] if d == "L" else path[key][1]
        i += 1
    print(i)
    return i


def part2(text) -> int:
    directions, path = parse(text)
    directions = cycle(directions)
    i = 0

    keys = [k for k in path if k[-1] == "A"]
    print("keys", keys)
    key_z_index = [0] * len(keys)
    while not all(k > 0 for k in key_z_index):
        d = next(directions)
        keys = [path[k][0] if d == "L" else path[k][1] for k in keys]
        i += 1
        for j, k in enumerate(keys):
            if k[-1] == "Z":
                key_z_index[j] = i
    print(key_z_index)
    result = lcm(*key_z_index)
    print(result)
    return result


if __name__ == "__main__":
    part1(StringIO(EXAMPLE_1))
    with open("input.txt") as f:
        part1(f)

    part2(StringIO(EXAMPLE_2))
    with open("input.txt") as f:
        part2(f)
