from dataclasses import dataclass
from io import StringIO
from typing import Iterable, List

EXAMPLE = """#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#"""

Pattern = List[str]


@dataclass
class Result:
    axis: str
    index: int

    @property
    def score(self):
        if self.axis == "v":
            return self.index
        return self.index * 100


def parse(text) -> Iterable[Pattern]:
    p = []
    for line in text.readlines():
        line = line.strip()
        if line == "":
            yield p
            p = []
        else:
            p.append(line)
    yield p


def part1(lefts, rights):
    return all(l[::-1] == r for l, r in zip(lefts, rights))


def part2(lefts, rights) -> int:
    """Find match score, which is count of number of chars
    which are not the exact mirror"""
    c = 0
    for l, r in zip(lefts, rights):
        for i, j in zip(l[::-1], r):
            if i != j:
                c += 1
    return c == 1


def find_mirror(pattern: Pattern, algo) -> Result:
    def _search(axis: str, p: Pattern):
        row_len = len(p[0])
        for i in range(1, row_len):
            s = min(i, row_len - i)
            lefts = [r[i - s : i] for r in p]
            rights = [r[i : i + s] for r in p]
            if algo(lefts, rights):
                return Result(axis=axis, index=i)
        return None

    result = _search("v", pattern)
    if result is None:
        # transpose
        row_len = len(pattern[0])
        pattern = ["".join([p[i] for p in pattern]) for i in range(row_len)]
        result = _search("h", pattern)
    if result is None:
        raise RuntimeError("oops")
    return result


def main(text, algo):
    results = []
    for p in parse(text):
        result = find_mirror(p, algo)
        results.append(result)
    print(results)
    score = sum(map(lambda r: r.score, results))
    print(score)


if __name__ == "__main__":
    main(StringIO(EXAMPLE), part1)
    with open("input.txt") as f:
        main(f, part1)

    main(StringIO(EXAMPLE), part2)
    with open("input.txt") as f:
        main(f, part2)
