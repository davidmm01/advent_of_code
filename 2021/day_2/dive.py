class Sub:
    def __init__(self):
        self.depth = 0
        self.position = 0
        self.aim = 0

    def go(self, direction, magnitude):
        fn = getattr(self, "go_" + direction)
        fn(magnitude)
        print(self.depth)
        print(self.position)
        print(self.aim)

    def final(self):
        return self.depth * self.position

    def go_forward(self, magnitude):
        raise NotImplemented

    def go_up(self, magnitude):
        raise NotImplemented

    def go_down(self, magnitude):
        raise NotImplemented


class Sub1(Sub):
    def go_forward(self, magnitude):
        self.position += magnitude

    def go_up(self, magnitude):
        self.depth -= magnitude

    def go_down(self, magnitude):
        self.depth += magnitude


class Sub2(Sub):
    def go_forward(self, magnitude):
        self.position += magnitude
        self.depth += self.aim * magnitude

    def go_up(self, magnitude):
        self.aim -= magnitude

    def go_down(self, magnitude):
        self.aim += magnitude


def get_input():
    with open("input.txt") as f:
        input_ = f.readlines()
    return [i.strip().split(" ") for i in input_]


def part_1():
    input_ = get_input()
    submarine = Sub1()
    for i in input_:
        submarine.go(i[0], int(i[1]))

    print("Part 1:", submarine.final())
    # solution == 1868935


def part_2():
    input_ = get_input()
    submarine = Sub2()

    for i in input_:
        submarine.go(i[0], int(i[1]))

    print("Part 2:", submarine.final())
    # solution == 1965970888


def main():
    part_1()
    part_2()


if __name__ == "__main__":
    main()
