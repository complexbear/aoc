from copy import deepcopy
from collections import deque, defaultdict
from dataclasses import dataclass
from io import StringIO
from typing import List, Union

EXAMPLE_1 = """..F7.
.FJ|.
SJ.L7
|F--J
LJ..."""

EXAMPLE_2 = """7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ"""

EXAMPLE_3 = """..........
.S------7.
.|F----7|.
.||....||.
.||....||.
.|L-7F-J|.
.|..||..|.
.L--JL--J.
.........."""

EXAMPLE_4 = """FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L"""


# Top left corner => (0, 0)
@dataclass(frozen=True)
class Point:
    x: int
    y: int


Path = List[Point]


class Maze:
    grid: List[List[str]] = None
    max_y = None
    max_x = None
    start = Point(0, 0)
    loop_path = None
    in_loop = []

    def __init__(self, text):
        lines = text.readlines()
        self.max_y = len(lines)
        self.max_x = len(lines[0]) - 1
        self.grid = [l.strip() for l in lines]
        self._find_start()

    def _find_start(self):
        for y, row in enumerate(self.grid):
            for x, cell in enumerate(row):
                if cell == "S":
                    self.start = Point(x, y)

    def _off_grid(self, p: Point) -> bool:
        return p.x < 0 or p.y < 0 or p.x >= self.max_x or p.y >= self.max_y

    @staticmethod
    def _update_point(p, x=0, y=0) -> Point:
        return Point(p.x + x, p.y + y)

    def _next_point(self, prev: Point, current: Point) -> Union[None, Point]:
        """Return next point from this start point"""
        next_ = Point(current.x, current.y)
        this = self.grid[current.y][current.x]
        if this == "|":
            next_ = self._update_point(next_, y=current.y - prev.y)
        elif this == "-":
            next_ = self._update_point(next_, x=current.x - prev.x)
        elif this == "L":
            next_ = self._update_point(
                next_,
                x=1 if current.x == prev.x else 0,
                y=-1 if current.y == prev.y else 0,
            )
        elif this == "7":
            next_ = self._update_point(
                next_,
                x=-1 if current.x == prev.x else 0,
                y=1 if current.y == prev.y else 0,
            )
        elif this == "J":
            next_ = self._update_point(
                next_,
                x=1 if current.x == prev.x else 0,
                y=-1 if current.y == prev.y else 0,
            )
        elif this == "F":
            next_ = self._update_point(
                next_,
                x=1 if current.x == prev.x else 0,
                y=1 if current.y == prev.y else 0,
            )

        if self._off_grid(next_):
            return None
        if self.grid[next_.y][next_.x] == ".":
            return None
        return next_ if self._valid_step(current, next_) else None

    def _valid_step(self, prev: Point, current: Point) -> bool:
        this = self.grid[current.y][current.x]
        last = self.grid[prev.y][prev.x]
        forbidden = None
        if last == "|":
            if current.y - prev.y > 0:
                forbidden = ("-", "F", "7")
            else:
                forbidden = ("-", "L", "J")
        elif last == "-":
            if current.x - prev.x > 0:
                forbidden = ("|", "L", "F")
            else:
                forbidden = ("|", "J", "7")
        elif last == "L":
            if current.x - prev.x > 0:
                forbidden = ("|", "L", "F")
            else:
                forbidden = ("-", "L", "J")
        elif last == "7":
            if current.x - prev.x < 0:
                forbidden = ("|", "7", "J")
            else:
                forbidden = ("-", "7", "F")
        elif last == "J":
            if current.x - prev.x < 0:
                forbidden = ("|", "J", "7")
            else:
                forbidden = ("-", "L", "J")
        elif last == "F":
            if current.x - prev.x > 0:
                forbidden = ("|", "F", "L")
            else:
                forbidden = ("-", "F", "7")
        return this not in forbidden

    def _path_step(self, path: Path) -> Union[None, Path]:
        current = path[-1]
        prev = path[-2]
        next_ = self._next_point(prev, current)
        if next_ is None or not self._valid_step(current, next_):
            return None
        return path + [next_]

    def find_loop(self):
        paths = deque()

        # add points around the start as possible paths
        for p in [
            Point(self.start.x + 1, self.start.y),
            Point(self.start.x - 1, self.start.y),
            Point(self.start.x, self.start.y + 1),
            Point(self.start.x, self.start.y - 1),
        ]:
            if self._off_grid(p):
                continue
            if self.grid[p.y][p.x] == ".":
                continue
            paths.append([self.start, p])

        while self.loop_path is None:
            path = paths.pop()
            while path is not None:
                if len(path) > 2 and path[-1] == self.start:
                    # loop complete
                    self.loop_path = path
                    break
                path = self._path_step(path)

    def print(self):
        output = deepcopy(self.grid)
        for p in self.loop_path:
            l = list(output[p.y])
            l[p.x] = "+"
            output[p.y] = "".join(l)
        for p in self.in_loop:
            l = list(output[p.y])
            l[p.x] = "I"
            output[p.y] = "".join(l)
        print("\n".join(output))

    def find_area(self):
        max_x = defaultdict(int)
        for p in self.loop_path:
            max_x[p.y] = max(max_x.get(p.y, 0), p.x)

        area_x = set()
        for y in range(self.max_y):
            inside = False
            for x in range(self.max_x):
                p = Point(x, y)
                if inside and p in self.loop_path:
                    inside = False
                if p in self.loop_path:
                    inside = True
                elif p.x >= max_x[p.y]:
                    inside = False
                elif inside:
                    area_x.add(p)
        area_y = set()
        for x in range(self.max_x):
            inside = False
            for y in range(self.max_y):
                p = Point(x, y)
                if inside and p in self.loop_path:
                    inside = False
                if p in self.loop_path:
                    inside = True
                elif p.x >= max_x[p.y]:
                    inside = False
                elif inside:
                    area_y.add(p)
        self.in_loop = area_x.intersection(area_y)


def main(text, find_area=False):
    m = Maze(text)
    print(m.max_y, m.max_x, m.start)
    m.find_loop()
    loop_size = len(m.loop_path)
    print(loop_size, (loop_size - 1) // 2)
    m.print()
    if find_area:
        m.find_area()
        m.print()
        print(len(m.in_loop))


if __name__ == "__main__":
    main(StringIO(EXAMPLE_1))
    main(StringIO(EXAMPLE_2))
    with open("input.txt") as f:
        main(f)

    main(StringIO(EXAMPLE_3), True)
    main(StringIO(EXAMPLE_4), True)

    with open("input.txt") as f:
        main(f, True)
