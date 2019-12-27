# Day 3, 2019 Advent of Code 

See [advent-of-code-day-3](https://adventofcode.com/2019/day/3)

## Part 1

### Plan of attack

1. Build an empty grid (a 2-dimensional array or matrix or something) of 0's.
2. Make it nice and big, probably like 10000 by 10000 or so
3. Initialise a cursor, starting in the middle at coordinates (5000, 5000)
4. build a translator that turns L100 into a move left by 100 units etc...
5. represent the path of the first wire by 1's in the grid
6. map out the second wire, and each time i piece of wire is placed, check if it intersects with the existing wire.  Record the intersection coordinates separately.  No need to record this mapping, since we don't care about self-intersection
7. Go through our recorded intersections and find the distance from each of the intersections to the central port.
