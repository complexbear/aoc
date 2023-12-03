from dataclasses import dataclass
from collections import defaultdict
from io import StringIO
from typing import List

EXAMPLE = """467..114..
    ...*......
    ..35..633.
    ......#...
    617*......
    .....+.58.
    ..592.....
    ......755.
    ...$.*....
    .664.598.."""


@dataclass(frozen=True)
class Coord:
    x: int
    y: int


@dataclass
class PartNum:
    value: int
    start: Coord
    end: Coord


def parse(text) -> (List[str], List[PartNum]):
    lines = text.readlines()
    numbers: List[PartNum] = []

    def store_num(num, x, y):
        numbers.append(
            PartNum(
                value=int(num),
                start=Coord(x - len(num), y=y),
                end=Coord(x=x - 1, y=y),
            )
        )

    for y, line in enumerate(lines):
        num = ""
        for x, c in enumerate(line.strip()):
            if c.isdigit():
                num += c
            elif num:
                store_num(num, x, y)
                num = ""
        if num:
            store_num(num, x, y)

    return list(map(str.strip, lines)), numbers


def is_adjacent(grid: List[str], number: PartNum) -> (bool, Coord):
    max_x = len(grid[0])
    max_y = len(grid)
    adj = False
    gear = None
    for x in range(max(number.start.x - 1, 0), min(number.end.x + 2, max_x)):
        for y in range(max(number.start.y - 1, 0), min(number.start.y + 2, max_y)):
            if grid[y][x] != "." and not grid[y][x].isdigit():
                adj = True
                if grid[y][x] == "*":
                    gear = Coord(x, y)
    return adj, gear


def main(text):
    grid, numbers = parse(text)
    # print(numbers)
    part_numbers = []
    gears = defaultdict(list)
    for num in numbers:
        adj, gear = is_adjacent(grid, num)
        if adj:
            part_numbers.append(num)
            gears[gear].append(num)

    # [print(n) for n in part_numbers]
    print(sum(map(lambda x: x.value, part_numbers)))

    gears = [nums[0].value * nums[1].value for nums in gears.values() if len(nums) == 2]
    print(f"sum gears: {sum(gears)}")


if __name__ == "__main__":
    main(StringIO(EXAMPLE))
    with open("input.txt") as f:
        main(f)
