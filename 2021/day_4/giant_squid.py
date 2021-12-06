class BingoSquare:
    # each bingo square is repesented by 1 flat array.  We can use indexes to check things.
    COLLECTIONS = [
        list(range(5)),  # row 1
        list(range(5,10)),  # row 2
        list(range(10,15)),  # row 3
        list(range(15,20)),  # row 4
        list(range(20,25)),  # row 5
        [0, 5, 10, 15, 20],  # column 1
        [1, 6, 11, 16, 21],  # column 2
        [2, 7, 12, 17, 22],  # column 3
        [3, 8, 13, 18, 23],  # column 4
        [4, 9, 14, 19, 24],  # column 5
    ]

    def __init__(self, lines):
        self.data = []
        for line in lines:
            # note extend not append, we want 1 long array
            self.data.extend(line.split())  # handles double space
        self.data = [int(i) for i in self.data]

    def process_draw(self, draw):
        # when a number is drawn, replace it with `None`
        self.data = [i if i!= draw else None for i in self.data]

    def is_win(self):
        for collection in self.COLLECTIONS:
            winner = True
            for i in collection:
                if self.data[i] != None:
                    winner = False
            if winner:
                return True
        return False

    def remaining_sum(self):
        remaining = [i for i in self.data if i is not None]
        return sum(remaining)


def get_input():
    with open("input.txt") as f:
        lines = f.readlines()

    # get the bingo draw numbers
    draws = lines[0].strip().split(",")
    draws = [int(i) for i in draws]  # covnert draws to ints

    bingo_squares = []
    i = 2
    while i <= len(lines):
        bingo_squares.append(BingoSquare(lines[i:i+5]))
        i += 6

    # draws is an array of strings
    # bingo_squares is an array of BingoSquare instances
    return draws, bingo_squares


def count_bingo_wins(squares):
    count = 0
    for i in squares:
        if i is None:
            count += 1
    return count


def part_1():
    draws, bingo_squares = get_input()
    for draw in draws:
        for bingo_square in bingo_squares:
            bingo_square.process_draw(draw)
            if bingo_square.is_win():

                print("Part 1:", bingo_square.remaining_sum() * draw)
                return
                # solution == 25410


def part_2():
    draws, bingo_squares = get_input()

    for draw in draws:

        for i in range(len(bingo_squares)):
            # dont bother with bongo squares that have already
            if bingo_squares[i] is None:
                continue

            bingo_squares[i].process_draw(draw)
            if bingo_squares[i].is_win():

                # stop at last win
                if count_bingo_wins(bingo_squares) == len(bingo_squares) - 1:
                    print("Part 2:", bingo_squares[i].remaining_sum() * draw)
                    return
                    # solution == 2730

                # if it wasn't the last win, keep going and ignore this bingo square from now on
                bingo_squares[i] = None

    # i think there would be much nicer ways to do this one, but ive already won lol


def main():
    part_1()
    part_2()


if __name__ == "__main__":
    main()
