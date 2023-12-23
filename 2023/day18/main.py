from dataclasses import dataclass
from io import StringIO
from typing import List

EXAMPLE = """R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)"""


Grid = List[str]


@dataclass
class Instruction:
    direction: str
    size: int
    colour: str


def parse(text) -> List[Instruction]:
    instructions = []
    for line in text.readlines():
        direction, size, colour = line.strip().split(" ")
        instructions.append(
            Instruction(direction=direction, size=int(size), colour=colour[1:-1])
        )
    return instructions


def dig(instructions: List[Instruction]) -> Grid:
    grid = ["#"]
    x, y = 0, 0
    min_x, max_x = 0, 0
    min_y, max_y = 0, 0
    for i in instructions:
        for _ in range(i.size):
            if i.direction == "R":
                pass
            elif i.direction == "U":
                pass
            elif i.direction == "D":
                pass
            elif i.direction == "L":
                pass
    return grid


def main(text):
    instructions = parse(text)
    print(instructions)


if __name__ == "__main__":
    main(StringIO(EXAMPLE))
