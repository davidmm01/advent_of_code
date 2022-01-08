def get_input(file_):
    with open(file_) as f:
        contents = f.readlines()
    output = []
    for i in contents:
        i = i.strip()
        line = []
        for j in i:
            line.append(int(j.strip()))
        output.append(line)
    return output


def validate_pos_row(val):
    if val < 0 or val > 99:
        return None
    return val


def validate_pos_col(val):
    if val < 0 or val > 99:
        return None
    return val


def get_coord_adjacents(row, col):
    # basic case, looking at row=2, col=6, 4 adjacent locations
    #       (1,6)
    # (2,5) (2,6) (2,7)
    #       (3,6)

    # edge case, looking at row=0, col=0, only 2 adjacent locations
    # (0,0) (0,1)
    # (1,0)

    row_low = validate_pos_row(row - 1)
    row_high = validate_pos_row(row + 1)
    col_low = validate_pos_col(col - 1)
    col_high = validate_pos_col(col + 1)

    adjacents = []

    if row_low is not None:
        adjacents.append([row_low, col])

    if row_high is not None:
        adjacents.append([row_high, col])

    if col_low is not None:
        adjacents.append([row, col_low])

    if col_high is not None:
        adjacents.append([row, col_high])

    return adjacents


def is_low_point(input_, row, col):
    point = input_[row][col]
    adjacents = get_coord_adjacents(row, col)
    for adjacent in adjacents:
        if point >= input_[adjacent[0]][adjacent[1]]:
            return False
    return True


def part_1():
    input_ = get_input("input.txt")
    # input_ is an array of arrays
    # reference by input_[row][column]

    # for each position, check if its a low point, and if it is, then sum the
    # risk level
    total_risk_level = 0
    for row in range(100):
        for col in range(100):
            if is_low_point(input_, row, col):
                total_risk_level += input_[row][col] + 1  # risk == 1 + height

    print("Part 1:", total_risk_level)
    # solution == 560


def part_2():
    # idea:
    # iterate through the input, and build up each entire basin.  Once a basin is built,
    # remember its size.
    # Remember which squares we have already considered by the `traversed_tracker`

    input_ = get_input("input.txt")
    rows = len(input_)
    cols = len(input_[0])
    global traversed_tracker
    traversed_tracker = [[False for i in range(cols)] for j in range(rows)]

    sizes = []
    for row in range(rows):
        for col in range(cols):
            global points_in_basin
            points_in_basin = []

            determine_inclusion_and_traverse(input_, row, col)

            if len(points_in_basin) != 0:
                sizes.append(len(points_in_basin))

    sizes.sort()

    print("Part 2:", sizes[-1] * sizes[-2] * sizes[-3])
    # solution == 959136


def determine_inclusion_and_traverse(input_, row, col):
    if input_[row][col] == 9:
        # not included in a basin
        traversed_tracker[row][col] = True
        return

    if traversed_tracker[row][col]:
        # already seen
        return

    # add to basin
    points_in_basin.append([row, col])
    traversed_tracker[row][col] = True

    # now go look at its adjacents
    adjacents = get_coord_adjacents(row, col)
    for adjacent in adjacents:
        determine_inclusion_and_traverse(input_, adjacent[0], adjacent[1])


def main():
    part_1()
    part_2()


if __name__ == "__main__":
    main()
