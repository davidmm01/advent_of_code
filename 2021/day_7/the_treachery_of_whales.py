def get_input():
    with open("input.txt") as f:
        line = f.readline()
    # '1,2,3' --> [1, 2, 3]
    return [int(i) for i in line.split(",")]


def make_tracker(input_):
     # determine the number of crabs at each horizontal position, stored in dict `tracker`
    tracker = {}
    for i in input_:
        if tracker.get(i) is None:
            tracker[i] = 1
        else:
            tracker[i] += 1
    return tracker


def part_1():
    input_ = get_input()
    # the optimal position will always be between (including) the min and max positions
    min_ = min(input_)
    max_ = max(input_)

    tracker = make_tracker(input_)

    best_total_fuel = None

    for candidate in range(min_, max_ + 1):
        # Note: here we are considering that each value between min_ and max_ + 1 is a potential
        # best case location.  I think this might be overkill, maybe you can get away with not checking
        # positions that don't already home a crab?

        # for each candidate location, figure out the total_fuel
        total_fuel = 0
        for location, quantity in tracker.items():

            fuel = 0
            if candidate >= location:
                fuel += candidate - location
            else:
                fuel += location - candidate
            fuel *= quantity

            total_fuel += fuel

        if best_total_fuel is None:
            best_total_fuel = total_fuel
        elif total_fuel < best_total_fuel:
            best_total_fuel = total_fuel

    print("Part 1:", best_total_fuel)
    # solution == 354129


def part_2():
    # Idea: part 2 is the same as part 1, except we must change the metric that calculates
    # how much fuel is used

    input_ = get_input()
    # the optimal position will always be between (including) the min and max positions
    min_ = min(input_)
    max_ = max(input_)
    tracker = make_tracker(input_)

    best_total_fuel = None

    for candidate in range(min_, max_ + 1):
        # for each candidate location, figure out the total_fuel
        total_fuel = 0
        for location, quantity in tracker.items():

            fuel = 0
            # sum of the first n natural numbers is: n(n+1) / 2
            if candidate >= location:
                n = candidate - location
            else:
                n = location - candidate

            fuel = int((n * (n + 1)) / 2)
            fuel *= quantity

            total_fuel += fuel

        if best_total_fuel is None:
            best_total_fuel = total_fuel
        elif total_fuel < best_total_fuel:
            best_total_fuel = total_fuel

    print("Part 2:", int(best_total_fuel))
    # solution == 98905973


def main():
    part_1()
    part_2()


if __name__ == "__main__":
    main()
