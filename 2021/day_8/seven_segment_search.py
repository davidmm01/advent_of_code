"""
  0:      1:      2:      3:      4:
 aaaa    ....    aaaa    aaaa    ....
b    c  .    c  .    c  .    c  b    c
b    c  .    c  .    c  .    c  b    c
 ....    ....    dddd    dddd    dddd
e    f  .    f  e    .  .    f  .    f
e    f  .    f  e    .  .    f  .    f
 gggg    ....    gggg    gggg    ....

  5:      6:      7:      8:      9:
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c
b    .  b    .  .    c  b    c  b    c
 dddd    dddd    ....    dddd    dddd
.    f  e    f  .    f  e    f  .    f
.    f  e    f  .    f  e    f  .    f
 gggg    gggg    ....    gggg    gggg

From this model, we will say:
a == row top
b == col top left
c == col top right
d == row middle
e == col bottom left
f == col bottom right
g == row bottom
"""

# Constants tracking the number of segments that each number uses
SEG_NO_0 = 6
SEG_NO_1 = 2
SEG_NO_2 = 5
SEG_NO_3 = 5
SEG_NO_4 = 4
SEG_NO_5 = 5
SEG_NO_6 = 6
SEG_NO_7 = 3
SEG_NO_8 = 7
SEG_NO_9 = 6


UNIQUES = (SEG_NO_1, SEG_NO_4, SEG_NO_7, SEG_NO_8)


class SevenSegment:
    def __init__(self, patterns):
        self.patterns = patterns
        self.row_top = None
        self.row_middle = None
        self.row_bottom = None
        self.col_top_left = None
        self.col_top_right = None
        self.col_bot_left = None
        self.col_bot_right = None

    def find_index(self, target_length):
        assert target_length in UNIQUES  # only use this for method uniques
        for i in range(len(self.patterns)):
            if len(self.patterns[i]) == target_length:
                return  i

    def determine_top_row(self):
        # can determine the `row_top` by seeing what segment differs for
        # representation of 1 (length 2) and 7 (length 3)
        index_of_1 = self.find_index(SEG_NO_1)
        index_of_7 = self.find_index(SEG_NO_7)
        for char in self.patterns[index_of_7]:
            if char not in self.patterns[index_of_1]:
                self.row_top = char
        assert self.row_top is not None


def get_input():
    input_ = []
    with open("input.txt") as f:
        lines = f.readlines()
    for line in lines:
        unique_patterns, output_values = line.split("|")
        unique_patterns = unique_patterns.split()
        output_values = output_values.split()
        input_.append(dict(patterns=unique_patterns, output=output_values))
    
    # returns [{patterns: [str, ...], output: [str, ...]}, ...]
    return input_


def part_1():
    input_ = get_input()
    count = 0
    for data in input_:
        for output_value in data["output"]:
            if len(output_value) in UNIQUES:
                count += 1
    print("Part 1:", count)
    # solution == 456


def part_2():
    input_ = get_input()


def main():
    part_1()
    part_2()


if __name__ == "__main__":
    main()
