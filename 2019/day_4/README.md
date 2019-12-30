# Day 4, 2019 Advent of Code 

See [advent-of-code-day-4](https://adventofcode.com/2019/day/4)

## Part 1

### Plan of attack

1. Make 3 functions that return true or false (depending on whether the number given satisfies the condition)
2. Feed through all the numbers that satisfy that condition.
3. Capture the numbers (this might be handy for part 2) and return how many there are

## Part 2

### Plan of attack

Modify our existing `has_double_digit()` to instead count the number of times each unique digit appears, and return whether or not a digit appears exactly twice or not.
