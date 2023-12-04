from dataclasses import dataclass, field
from collections import deque, defaultdict

from io import StringIO
from math import pow
from typing import List, Set

EXAMPLE = """Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"""


@dataclass
class Card:
    winning: Set[int]
    chosen: Set[int]
    matched: Set[int] = field(init=False)

    def __post_init__(self):
        self.matched = self.chosen.intersection(self.winning)


def parse(text) -> List[Card]:
    cards = []
    for line in text.readlines():
        winning, chosen = [
            set(map(int, s.strip().split())) for s in line.split(":")[1].split("|")
        ]
        cards.append(Card(winning=winning, chosen=chosen))
    return cards


def main(text):
    sum_cards = 0
    cards = parse(text)

    # part 1
    for c in cards:
        if c.matched:
            score = int(pow(2, len(c.matched) - 1))
            print(score)
            sum_cards += score
    print(f"part 1 sum cards:", sum_cards)

    # part 2
    # map card number to which cards are won by it
    originals = {
        idx: list(range(idx + 1, idx + 1 + len(c.matched)))
        for idx, c in enumerate(cards)
    }

    totals = defaultdict(int)
    q = deque(range(len(cards)))
    while q:
        i = q.popleft()
        totals[i] += 1
        q.extend(originals[i])
    print(totals)
    print(f"part 2 sum cards: {sum(totals.values())}")


if __name__ == "__main__":
    main(StringIO(EXAMPLE))
    with open("input.txt") as f:
        main(f)
