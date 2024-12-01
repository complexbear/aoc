from collections import namedtuple
from io import StringIO
from typing import List, Set

EXAMPLE = """...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
..........."""

Point = namedtuple("Point", "x, y")
Grid = List[str]


def parse(text) -> (Point, Grid):
    start = None
    grid = [line.strip() for line in text.readlines()]
    for y in range(len(grid)):
        x = grid[y].index("S") if "S" in grid[y] else None
        if x is not None:
            start = Point(x, y)
    return start, grid


def step(grid: Grid, plots: Set[Point]) -> Set[Point]:
    max_x = len(grid[0])
    max_y = len(grid)
    new_plots = set([])

    def _next_plot(p: Point):
        x = p.x % max_x
        y = p.y % max_y
        # jump to next grid
        if p.x < 0:
            x = max_x
        if p.y < 0:
            y = max_y - y
        try:
            if grid[y][x] == "#":
                return
        except Exception:
            print("doh")
        new_plots.add(p)

    while plots:
        p = plots.pop()
        for dx in [-1, 1]:
            _next_plot(Point(p.x + dx, p.y))
        for dy in [-1, 1]:
            _next_plot(Point(p.x, p.y + dy))
    return new_plots


def main(text, steps: int):
    start, grid = parse(text)
    # print(start)
    plots = {start}
    for _ in range(steps):
        plots = step(grid, plots)
    print(steps, len(plots))


if __name__ == "__main__":
    for s in (6, 10, 50, 100, 1000, 5000):
        main(StringIO(EXAMPLE), s)

    # with open("input.txt") as f:
    #     main(f, 64)
