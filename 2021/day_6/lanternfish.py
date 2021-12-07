class LaternfishSchool:
    def __init__(self, fish_spawn_timers):
        self.timers = {
            0: 0,
            1: 0,
            2: 0,
            3: 0,
            4: 0,
            5: 0,
            6: 0,
            7: 0,
            8: 0,
        }
        for i in fish_spawn_timers:
            self.timers[i] += 1

    def pass_day(self):
        # note here we will store new values in `new_vals`, since we don't want to update the
        # `self.timers` dict while we are still iterating through it!
        new_vals = {
            0: 0,
            1: 0,
            2: 0,
            3: 0,
            4: 0,
            5: 0,
            6: 0,
            7: 0,
            8: 0,
        }
        for timer, number in self.timers.items():
            if timer == 0:
                # spawn time: all these fish will return to timer 6, and that many new fish will
                # go to timer 8
                new_vals[8] += number
                new_vals[6] += number
            else:
                new_vals[timer - 1] += number

        self.timers = new_vals

    def count_all_fish(self):
        return sum(self.timers.values())


def get_input():
    with open("input.txt") as f:
        line = f.readline()
    # '1,2,3' --> [1, 2, 3]
    return [int(i) for i in line.split(",")]


def part_1():
    input_ = get_input()
    school = LaternfishSchool(input_)
    for i in range(80):
        school.pass_day()
    print("Part 1:", school.count_all_fish())
    # solution == 350605


def part_2():
    input_ = get_input()
    school = LaternfishSchool(input_)
    for i in range(256):
        school.pass_day()
    print("Part 2:", school.count_all_fish())
    # solution == 1592778185024
    # so easy... i guess part 2 punishes people who don't do part 1 efficiently?


def main():
    part_1()
    part_2()


if __name__ == "__main__":
    main()
