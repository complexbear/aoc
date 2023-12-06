from io import StringIO
from functools import reduce
import operator

EXAMPLE = """Time:      7  15   30
Distance:  9  40  200"""


def part1(time, record):
    return len([h for h in range(time + 1) if h * (time - h) > record])


def part2(time, record):
    first_win_idx = None
    for h in range(time + 1):
        if h * (time - h) > record:
            first_win_idx = h
            break
    ways_to_win = time - (2 * first_win_idx - 1)
    return ways_to_win


def parse(text, algo):
    lines = text.readlines()
    if algo is part1:
        data = [list(map(int, l.split(":")[1].strip().split())) for l in lines]
        return zip(*data)
    else:
        time, record = [int(l.split(":")[1].replace(" ", "")) for l in lines]
        return [(time, record)]


def main(text, algo):
    races = parse(text, algo)
    # dist = hold * (time-hold)
    ways_to_win = [algo(t, r) for t, r in races]
    print(ways_to_win)
    print(reduce(operator.mul, ways_to_win))


if __name__ == "__main__":
    main(StringIO(EXAMPLE), part1)
    with open("input.txt") as f:
        main(f, part1)

    print("part 2")
    main(StringIO(EXAMPLE), part2)
    with open("input.txt") as f:
        main(f, part2)
