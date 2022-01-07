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

# Uniques:
# Number 1 w 2 segments
# Number 4 w 4 segments
# Number 7 w 3 segments
# Number 8 w 7 segments

UNIQUES = (SEG_NO_1, SEG_NO_4, SEG_NO_7, SEG_NO_8)

# letters used to represent segments
SEG_CHARS = ["a", "b", "c", "d", "e", "f", "g"]

class SevenSegment:
    def __init__(self, patterns, outputs):
        self.patterns = patterns
        self.outputs = outputs
        self.output_number = ""
        self.row_top = None
        self.row_middle = None
        self.row_bottom = None
        self.col_top_left = None
        self.col_top_right = None
        self.col_bot_left = None
        self.col_bot_right = None

        # self.pretty_print()
        self._determine_row_top()
        # self.pretty_print()
        self._determine_col_right()
        # self.pretty_print()
        self._determine_top_left_and_row_mid()
        # self.pretty_print()
        self._determine_row_bottom_and_col_bot_left()
        # self.pretty_print()

        for output in self.outputs:
            self.output_number = self.output_number + self.determine_number(output)

    def find_index(self, target_length):
        assert target_length in UNIQUES  # only use this for method uniques
        for i in range(len(self.patterns)):
            if len(self.patterns[i]) == target_length:
                return  i

    def _determine_row_top(self):
        # can determine the `row_top` by seeing what segment differs for
        # representation of 1 (length 2) and 7 (length 3)
        index_of_1 = self.find_index(SEG_NO_1)
        index_of_7 = self.find_index(SEG_NO_7)
        for char in self.patterns[index_of_7]:
            if char not in self.patterns[index_of_1]:
                self.row_top = char
        assert self.row_top is not None

    def _determine_col_right(self):
        # From the represnetation of 1, we know the two segments that make up the
        # right column. But which is which? 
        # We can determine by looking at all the numbers with 6 segments, those
        # being 0, 6 and 9, since all 3 will have `self.col_bot_right`, but only 2 of them
        # will have `self.col_top_right`.
        val_1, val_2 = self.patterns[self.find_index(SEG_NO_1)]
        val_1_seen, val_2_seen = 0, 0
        for pattern in self.patterns:
            if len(pattern) == 6:
                if val_1 in pattern:
                    val_1_seen += 1
                if val_2 in pattern:
                    val_2_seen += 1

        assert 2 in (val_1_seen, val_2_seen)
        assert 3 in (val_1_seen, val_2_seen)

        if val_1_seen == 2:
            self.col_top_right = val_1
            self.col_bot_right = val_2
        else:
            self.col_top_right = val_2
            self.col_bot_right = val_1

    def _determine_top_left_and_row_mid(self):
        # Now we know top row and the right column, we can use this when looking at
        # number 4 (only number with 4 segments) and determine the 2 segments that can
        # be `self.row_middle` and `self.col_top_right`. Next we can again look at
        # the numbers with 6 segments, and determine which are which. Each of
        # 0, 6 and 9 will have the col top right, but only 2 of them will have row_middle
        vals = self.patterns[self.find_index(SEG_NO_4)]
        val_1 = None
        val_2 = None
        for val in vals:
            if val not in (self.col_top_right, self.col_bot_right):
                if val_1 is None:
                    val_1 = val
                else:
                    val_2 = val
        assert val_1 is not None
        assert val_2 is not None
        
        val_1_seen, val_2_seen = 0, 0

        for pattern in self.patterns:
            if len(pattern) == 6:
                if val_1 in pattern:
                    val_1_seen += 1
                if val_2 in pattern:
                    val_2_seen += 1

        assert 2 in (val_1_seen, val_2_seen)
        assert 3 in (val_1_seen, val_2_seen)

        if val_1_seen == 2:
            self.row_middle = val_1
            self.col_top_left = val_2
        else:
            self.row_middle = val_2
            self.col_top_left = val_1

    def _determine_row_bottom_and_col_bot_left(self):
        # determine the two remaining segments, that will be the candidates for column bot left and row bottom.
        val_1, val_2 = set([self.row_top, self.row_middle, self.col_top_left, self.col_top_right, self.col_bot_right]) ^ set(SEG_CHARS)
        val_1_seen, val_2_seen = 0, 0

        # continue doing the same kind of logic we have been doing, but this time with 5 segment numbers and searching
        # for appearances 1 and 3.
        for pattern in self.patterns:
            if len(pattern) == 5:
                if val_1 in pattern:
                    val_1_seen += 1
                if val_2 in pattern:
                    val_2_seen += 1

        assert 1 in (val_1_seen, val_2_seen)
        assert 3 in (val_1_seen, val_2_seen)

        if val_1_seen == 1:
            self.col_bot_left = val_1
            self.row_bottom = val_2
        else:
            self.col_bot_left = val_2
            self.row_bottom = val_1

    def _is_0(self, pattern):
        return len(pattern) == 6 and self.row_top in pattern and self.row_bottom in pattern and self.col_top_left in pattern and self.col_top_right in pattern and self.col_bot_left in pattern and self.col_bot_right in pattern

    def _is_1(self, pattern):
        return len(pattern) == 2

    def _is_2(self, pattern):
        return len(pattern) == 5 and self.row_top in pattern and self.row_middle in pattern and self.row_bottom in pattern and self.col_top_right in pattern and self.col_bot_left in pattern

    def _is_3(self, pattern):
        return len(pattern) == 5 and self.row_top in pattern and self.row_middle in pattern and self.row_bottom in pattern and self.col_bot_right in pattern and self.col_top_right in pattern

    def _is_4(self, pattern):
        return len(pattern) == 4

    def _is_5(self, pattern):
        return len(pattern) == 5 and self.row_top in pattern and self.row_middle in pattern and self.row_bottom in pattern and self.col_top_left in pattern and self.col_bot_right in pattern

    def _is_6(self, pattern):
        return len(pattern) == 6 and self.row_top in pattern and self.row_middle in pattern and self.row_bottom in pattern and self.col_top_left in pattern and self.col_bot_left in pattern and self.col_bot_right in pattern

    def _is_7(self, pattern):
        return len(pattern) == 3

    def _is_8(self, pattern):
        return len(pattern) == 7

    def _is_9(self, pattern):
        return len(pattern) == 6 and self.row_top in pattern and self.row_middle in pattern and self.row_bottom in pattern and self.col_top_left in pattern and self.col_top_right in pattern and self.col_bot_right in pattern

    def determine_number(self, pattern):
        if self._is_0(pattern):
            return "0"
        if self._is_1(pattern):
            return "1"
        if self._is_2(pattern):
            return "2"
        if self._is_3(pattern):
            return "3"
        if self._is_4(pattern):
            return "4"
        if self._is_5(pattern):
            return "5"
        if self._is_6(pattern):
            return "6"
        if self._is_7(pattern):
            return "7"
        if self._is_8(pattern):
            return "8"
        if self._is_9(pattern):
            return "9"

    def pretty_print(self):
        # only required for development/debug

        def to_dot(value):
            if value is None:
                return "."
            return value

        print(f"\n {to_dot(self.row_top)}{to_dot(self.row_top)}{to_dot(self.row_top)} ")
        print(f"{to_dot(self.col_top_left)}   {to_dot(self.col_top_right)}")
        print(f"{to_dot(self.col_top_left)}   {to_dot(self.col_top_right)}")
        print(f" {to_dot(self.row_middle)}{to_dot(self.row_middle)}{to_dot(self.row_middle)} ")
        print(f"{to_dot(self.col_bot_left)}   {to_dot(self.col_bot_right)}")
        print(f"{to_dot(self.col_bot_left)}   {to_dot(self.col_bot_right)}")
        print(f" {to_dot(self.row_bottom)}{to_dot(self.row_bottom)}{to_dot(self.row_bottom)} ")

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
    inputs = get_input()
    total = 0
    for input_ in inputs:
        # all the hard work to determine the number is done in the SevenSegment init
        seven_segment = SevenSegment(input_["patterns"], input_["output"])
        total += int(seven_segment.output_number)
    print("Part 2:", total)
    # solution == 1091609

def main():
    part_1()
    part_2()


if __name__ == "__main__":
    main()
