from collections import namedtuple
from functools import partial, reduce
from io import StringIO
from typing import List

EXAMPLE = """seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4"""

Range = namedtuple("Range", "dst, src, size")
Map = List[Range]
Maps = List[Map]


def text_to_ints(line):
    return list(map(int, line.split()))


def parse(text) -> (List[int], Maps):
    seeds = text_to_ints(text.readline().split(":")[1])
    maps = []
    ranges = []
    for line in text.readlines():
        if "map:" in line:
            if ranges:
                maps.append(ranges)
            ranges = []
        elif line.strip() == "":
            continue
        else:
            dst, src, size = text_to_ints(line)
            ranges.append(Range(dst=dst, src=src, size=size))
    if ranges:
        maps.append(ranges)
    return seeds, maps


def apply_map(s: int, r: Range) -> (bool, int):
    if r.src <= s <= r.src + r.size:
        offset = s - r.src
        return True, r.dst + offset
    return False, s


def main(text, debug: bool):
    seeds, maps = parse(text)

    if debug:
        print(f"seeds: {seeds}")
        for i, m in enumerate(maps):
            print(f"{i}: {m}")

    locations = []
    for s in seeds:
        x = [s]
        for m in maps:
            y = x[-1]
            for r in m:
                mapped, y = apply_map(y, r)
                if mapped:
                    break
            x.append(y)
        if debug:
            print(f"{s}: {x}")
        locations.append(x[-1])
    print(f"locations: {locations}")
    print(min(locations))


if __name__ == "__main__":
    main(StringIO(EXAMPLE), True)
    with open("input.txt") as f:
        main(f, False)
