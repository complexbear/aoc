import operator
from functools import reduce
from io import StringIO
from enum import Enum
from typing import Dict, Iterable, List


class Colour(Enum):
    RED = "red"
    BLUE = "blue"
    GREEN = "green"


Sample = Dict[Colour, int]
Game = List[Sample]
Bag = Sample

EXAMPLE_1 = """Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"""


def parse(text) -> List[Game]:
    games = []
    lines = text.readlines()
    for line in lines:
        game = []
        samples = line.split(":")[1].split(";")
        for sample in samples:
            items = [s.strip().split() for s in sample.split(",")]
            game.append({Colour(i[1]): int(i[0]) for i in items})
        games.append(game)
    return games


def is_possible(bag: Bag, game: Game) -> (bool, Sample):
    valid = True
    max_colours = {Colour.RED: 0, Colour.BLUE: 0, Colour.GREEN: 0}
    for sample in game:
        for colour, count in sample.items():
            max_colours[colour] = max(max_colours[colour], count)
            if bag.get(colour, 0) < count:
                valid = False
    return valid, max_colours


def main(text):
    games = parse(text)

    bag = {Colour.RED: 12, Colour.BLUE: 14, Colour.GREEN: 13}
    fewest_colours = {}
    valid_games = []
    for idx, game in enumerate(games):
        valid, max_colours = is_possible(bag, game)
        fewest_colours[idx + 1] = max_colours
        if valid:
            valid_games.append(idx + 1)
    return valid_games, sum(valid_games), fewest_colours


def cube_sum(samples: Iterable[Sample]) -> int:
    sum_ = 0
    for sample in samples:
        cube = reduce(operator.mul, sample.values())
        print(f"cube: {cube}")
        sum_ += cube
    return sum_


if __name__ == "__main__":
    valid_games, answer, fewest_colours = main(StringIO(EXAMPLE_1))
    print(f"Example 1: {valid_games}, {answer}, {fewest_colours}")
    print(f"Example 2: {cube_sum(fewest_colours.values())}")

    with open("input.txt") as f:
        _, answer, fewest_colours = main(f)
        print(f"Part 1: {answer}")
        print(f"Part 2: {cube_sum(fewest_colours.values())}")
