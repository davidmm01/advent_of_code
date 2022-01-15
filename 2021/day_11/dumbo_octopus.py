def get_input():
    # get the input, where each line of the input will be represented as a list of integers
    lines = []
    with open("input.txt") as f:
        c = f.readline()
        while c:
            c = c.strip()
            line = []
            for i in c:
                line.append(int(i))
            lines.append(line)
            c = f.readline()
    return lines


ROWS = 10
COLS = 10
ROW_IDX_MIN = 0
ROW_IDX_MAX = ROWS - 1
COL_IDX_MIN = 0
COL_IDX_MAX = COLS - 1
STEPS = 100


class Octopus:
    def __init__(self, row, col, starting_energy):
        self.row = row
        self.col = col
        self.energy = starting_energy
        self.flashes = 0
        self.currently_flashing = False
        # the surrounding coords for each octopus won't change, so calculate it once on init of octopus
        self.surrounding = self.get_surrounding_coords()

    def increase_energy_and_notify_flashing(self):
        if self.energy < 9:
            self.energy += 1
            return False

        elif self.energy == 9 and not self.currently_flashing:
            self.energy += 1
            self.flashes += 1
            self.currently_flashing = True
            return True

        return False

    def post_step_reset(self):
        if self.energy > 9:
            self.energy = 0
            self.currently_flashing = False

    def get_surrounding_coords(self):
        # for this row and col pair, get all the surrounding pairs of [row, col]

        coords = [
            [self.row, self.col - 1],
            [self.row, self.col + 1],
            [self.row - 1, self.col],
            [self.row + 1, self.col],
            [self.row - 1, self.col - 1],
            [self.row - 1, self.col + 1],
            [self.row + 1, self.col - 1],
            [self.row + 1, self.col + 1],
        ]

        valid_coords = []
        for coord in coords:
            if self.is_coord_valid(coord):
                valid_coords.append(coord)

        return valid_coords

    @staticmethod
    def is_coord_valid(coord):
        row = coord[0]
        col = coord[1]
        return (
            row >= ROW_IDX_MIN
            and row <= ROW_IDX_MAX
            and col >= COL_IDX_MIN
            and col <= COL_IDX_MAX
        )


def print_energy_level(octos):
    string_rep = ""
    for row in octos:
        for octopus in row:
            string_rep += str(octopus.energy)
        string_rep += "\n"
    print(string_rep)


def part_1():
    # create the 2d array of octopusses
    input_ = get_input()
    octos = [[None for i in range(COLS)] for j in range(ROWS)]
    for row in range(ROWS):
        for col in range(COLS):
            octos[row][col] = Octopus(row, col, input_[row][col])

    for i in range(STEPS):
        to_flash = []
        # in each step:
        # each ocotopus has its energy level increased
        for row in octos:
            for octopus in row:
                octopus.increase_energy_and_notify_flashing()
                if octopus.currently_flashing:
                    # record octopuses that are flashing...
                    to_flash.append(octopus)

        # Now while there are octopus that are flashing, increase the energy of the surrounding octopus
        # while also adding any chain reaction flashes to the to_flash array.  Keep going until there is nothing
        # left to flash.
        while len(to_flash) > 0:
            current = to_flash[0]
            del to_flash[0]
            for coord in current.surrounding:
                if octos[coord[0]][coord[1]].increase_energy_and_notify_flashing():
                    to_flash.append(octos[coord[0]][coord[1]])

        # reset all the octopuses
        for row in octos:
            for octopus in row:
                octopus.post_step_reset()

        # debug only
        #  print_energy_level(octos)

    flashes = 0
    for row in octos:
        for octopus in row:
            flashes += octopus.flashes

    print("Part 1:", flashes)
    # solution == 1755


def part_2():
    # create the 2d array of octopusses
    input_ = get_input()
    octos = [[None for i in range(COLS)] for j in range(ROWS)]
    for row in range(ROWS):
        for col in range(COLS):
            octos[row][col] = Octopus(row, col, input_[row][col])

    counter = 0
    while True:
        flashing = 0
        counter += 1
        to_flash = []
        # in each step:
        # each ocotopus has its energy level increased
        for row in octos:
            for octopus in row:
                octopus.increase_energy_and_notify_flashing()
                if octopus.currently_flashing:
                    # record octopuses that are flashing...
                    to_flash.append(octopus)

        # Now while there are octopus that are flashing, increase the energy of the surrounding octopus
        # while also adding any chain reaction flashes to the to_flash array.  Keep going until there is nothing
        # left to flash.
        while len(to_flash) > 0:
            current = to_flash[0]
            del to_flash[0]
            for coord in current.surrounding:
                if octos[coord[0]][coord[1]].increase_energy_and_notify_flashing():
                    to_flash.append(octos[coord[0]][coord[1]])

        # check to see if every octopus flashed this step, if yes then we are done
        for row in octos:
            for octopus in row:
                if octopus.currently_flashing:
                    flashing += 1
        if flashing == 100:
            break

        for row in octos:
            for octopus in row:
                octopus.post_step_reset()

        # debug only
        #  print_energy_level(octos)

    print("Part 2:", counter)
    # solution == 212


def main():
    part_1()
    part_2()


if __name__ == "__main__":
    main()
