# Day 1, 2019 Advent of Code 

See [advent-of-code-day-1](https://adventofcode.com/2019/day/1)

## PART 1

### TASK:
Given a list of module masses, find the total amount of fuel required for all the modules

To find the fuel required, we must:
1. take the mass
2. divide by three
3. down down
4. subtract 2

Examples
- for a mass of 12, `floor(12/3) - 2 = 2`
- for a mass of 14, `floor(14/3) - 2 = 2`
- for a mass of 1969, `floor(1969/3) - 2 = 654`
- for a mass of 100756, `floor(100756/3) -2 = 33583`

### INPUT:
`input.txt`, a newline separated text file with each line being the mass of a module


## PART 2:

### TASK: 
Fuel has weight, so fuel also needs fuel.  Once you have found the required fuel for a module, find out the fuel that fuel needs.  
Keep iterating until each modules' fuel is fueled sufficiently.
If a fuel needs 0 or negative fuel, then we can assume the fuel has been accounted for sufficiently.

Example
- for a mass of 1969, `floor(1969/3) - 2 = 654`
- then for the fuel mass of 654, `floor(654/3) - 2 = 216`
- then for the fuels fuel of mass 216, `floor(216/3) - 2 = 72`
- then for the fuels fuels fuel of mass 72, `floor(72/3) - 2 = 22`
- then for the fuels fuels fuels fuel of mass 24, `floor(24/3) - 2 = 6`
- then for the fuels fuels fuels fuels fuel of mass 6, `floor(6/3) - 2 = 0`
- so the total fuel requirement is: `total = 654 + 216 + 72 + 22 + 6 = 970`

### INPUT:
Exact same as part 1, `input.txt`
