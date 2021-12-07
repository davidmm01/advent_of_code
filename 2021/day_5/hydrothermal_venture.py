class CoordPair:
    def __init__(self, x1, y1, x2, y2):
        self.x1 = x1
        self.y1 = y1
        self.x2 = x2
        self.y2 = y2

    def is_vert_or_hori_line(self):
        return self.x1 == self.x2 or self.y1 == self.y2

    def make_vert_or_hori_line(self):
        line = []

        # horizontal
        if self.x1 == self.x2:
            y_low = min(self.y1, self.y2)
            y_high = max(self.y1, self.y2)
            for i in range(y_low, y_high+1):
                line.append((self.x1, i))

        # vertical
        elif self.y1 == self.y2:
            x_low = min(self.x1, self.x2)
            x_high = max(self.x1, self.x2)
            for i in range(x_low, x_high+1):
                line.append((i, self.y1))

        return line

    def make_vert_hori_diag_line(self):
        # basic line from before
        if self.is_vert_or_hori_line():
            return self.make_vert_or_hori_line()

        # else its a diagonal line
        assert abs(self.y2 - self.y1) == abs(self.x2 - self.x1)
        # THIS ONLY WORKS FOR 45 DEGREE LINES AS PER THE PUZZLE!

        # determine the x values in the line
        x_low = min(self.x1, self.x2)
        x_high = max(self.x1, self.x2)
        x_vals = list(range(x_low, x_high + 1))

        # switch the x values if required
        if self.x1 == x_high:
            x_vals.reverse()

        # determine the y values in the line
        y_low = min(self.y1, self.y2)
        y_high = max(self.y1, self.y2)
        y_vals = list(range(y_low, y_high + 1))

        # switch the y values if required
        if self.y1 == y_high:
            y_vals.reverse()

        # stick the x and y together
        line = []
        for x, y in zip(x_vals, y_vals):
            line.append([x, y])

        return line


class Map:
    def __init__(self):
        self.data = []
        for i in range(1000):
            self.data.append([0] * 1000)
        # self.data[x][y] == (x,y) coord

    def saturated_points(self, threshold):
        count = 0
        for row in self.data:
            for point in row:
                if point >= threshold:
                    count += 1
        return count

    def write_points(self, line):
        for coord in line:
            self.data[coord[0]][coord[1]] += 1


def get_input():
    with open("input.txt") as f:
        lines = f.readlines()
    lines = [line.strip().replace(" -> ", ",").split(",") for line in lines]
    coord_pairs = []
    for line in lines:
        coord_pairs.append(CoordPair(int(line[0]), int(line[1]), int(line[2]), int(line[3])))
    return coord_pairs


def part_1():
    coord_pairs = get_input()
    map_ = Map()
    for coord_pair in coord_pairs:
        if coord_pair.is_vert_or_hori_line():
            map_.write_points(coord_pair.make_vert_or_hori_line())
    print("Part 1:", map_.saturated_points(2))
    # solution == 5774


def part_2():
    coord_pairs = get_input()
    map_ = Map()
    for coord_pair in coord_pairs:
        map_.write_points(coord_pair.make_vert_hori_diag_line())
    print("Part 2:", map_.saturated_points(2))
    # solution == 18423


def main():
    part_1()
    part_2()


if __name__ == "__main__":
    main()
