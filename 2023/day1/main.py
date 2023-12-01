from io import StringIO

EXAMPLE_1 = """1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet"""

EXAMPLE_2 = """two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen"""

NUMBERS = [
    "zero",
    "one",
    "two",
    "three",
    "four",
    "five",
    "six",
    "seven",
    "eight",
    "nine",
]
MAX_NUMBER_WORD = 5


def findWord(line):
    for n, number in enumerate(NUMBERS):
        idx = line[:MAX_NUMBER_WORD].find(number)
        if idx == 0:
            return n
    return None


def extractNumber(line: str, with_word: bool):
    a, b = None, None
    i, j = 0, len(line) - 1
    while a is None or b is None:
        if a is None:
            if line[i].isdigit():
                a = line[i]
            if with_word:
                n = findWord(line[i:])
                if n is not None:
                    a = n

        if b is None:
            if line[j].isdigit():
                b = line[j]
            if with_word:
                n = findWord(line[j:])
                if n is not None:
                    b = n

        i += 1
        j -= 1
    return (int(a) * 10) + int(b)


def solution(text, part: int):
    data = text.readlines()
    numbers = [extractNumber(line, with_word=part == 2) for line in data]
    return numbers


if __name__ == "__main__":
    numbers = solution(StringIO(EXAMPLE_1), 1)
    print(f"Example: {numbers}, {sum(numbers)}")

    with open("input.txt") as f:
        numbers = solution(f, 1)
        print(f"Solution: {sum(numbers)}")

    numbers = solution(StringIO(EXAMPLE_2), 2)
    print(f"Example: {numbers}, {sum(numbers)}")

    with open("input.txt") as f:
        numbers = solution(f, 2)
        print(f"Solution: {sum(numbers)}")
