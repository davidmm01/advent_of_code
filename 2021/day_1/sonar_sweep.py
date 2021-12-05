def get_input():
    with open("input.txt") as f:
        input_ = f.readlines()
    return [int(i.strip()) for i in input_]


def part_1():
    input_ = get_input()
    previous = None
    increase_counter = 0

    for i in input_:
        # handle initial case
        if previous is None:
            previous = i
            continue

        if i > previous:
            increase_counter += 1

        previous = i

    print("Part 1:", increase_counter)
    # solution == 1581


def part_2():
    input_ = get_input()
    previous = None
    increase_counter = 0

    for i in range(len(input_) - 2):
        total = input_[i] + input_[i + 1] + input_[i + 2]

        # handle initial case
        if previous is None:
            previous = total
            continue

        if total > previous:
            increase_counter += 1

        previous = total

    print("Part 2:", increase_counter)
    # solution == 1996


def main():
    part_1()
    part_2()


if __name__ == "__main__":
    main()
