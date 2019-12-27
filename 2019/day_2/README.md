# Day 2, 2019 Advent of Code 

See [advent-of-code-day-2](https://adventofcode.com/2019/day/2)

## PART 1

### TASK/APPROACH:
Given a list of numbers, apply the various opcode protocols to find out what number remains at position 0 upon completion 

Opcode  1: Add protocol
Opcode  2: Multiply protocol
Opcode 99: Program complete

0. Think of our input as an array `input`.
1. replace input[1] with 12 and input[2] with 2
2. Let cursor = 0
3. If the first element `input[cursor + 0]` is 1 or 2, we perform the add or multiply protocol, else if its 99 we are done.  
4. Take elements `input[input[cursor + 1]]` and `input[input[cursor + 2]]`, add or multiply them (depending on protocol)
5. Store the result of the protocol in `input[cursor + 3]`
6. Cursor += 4
7. Repeat from step 2

### INPUT:
`input.txt`, a one text file with comma separated values


## PART 2

### TASK/APPROACH:

Just see the link for this one, essentially a little brute force to work backwards and find the input for a certain output (changing the values that get subbed into input[1] and input[2])
