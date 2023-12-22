from dataclasses import dataclass
from functools import reduce
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
    index: str

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


def find_mirror(pattern: Pattern) -> Result:
    def _search(axis: str, p: Pattern):
        row_len = len(p[0])
        for i in range(1, row_len):
            s = min(i, row_len - i)
            lefts = [r[i - s : i] for r in p]
            rights = [r[i : i + s] for r in p]
            if all(l[::-1] == r for l, r in zip(lefts, rights)):
                return Result(axis=axis, index=i)

    result = _search("v", pattern)
    if result is None:
        # transpose
        row_len = len(pattern[0])
        pattern = ["".join([p[i] for p in pattern]) for i in range(row_len)]
        result = _search("h", pattern)
    if result is None:
        raise RuntimeError("oops")
    return result


def main(text):
    results = []
    for p in parse(text):
        result = find_mirror(p)
        results.append(result)
    print(results)
    score = sum(map(lambda r: r.score, results))
    print(score)


if __name__ == "__main__":
    # main(StringIO(EXAMPLE))
    with open("input.txt") as f:
        main(f)
