def get_input():
    with open("input.txt") as f:
        input_ = f.readlines()
    return [i.strip() for i in input_]


def part_1():
    input_ = get_input()

    max_ = len(input_)
    # max_ is the greatest possible number of 1s or 0s for each bit
    threshold = max_ / 2
    # threshold is the number of 1s or 0s needed to win majority

    bits = len(input_[0])
    # bits is the number of bits in each binary number

    binary_tracker = [0] * bits
    # binary_tracker shows the number of 1's for each bit.
    # So binary_tracker[0] counts all the 1s seen for the first bit of each binary number.
    # From binary tracker we can infer that if the number of 1's is greater
    # than threshold, then the gamma rate bit is 1, else 0

    for binary_number in input_:
        for pos, bit in enumerate(binary_number):
            binary_tracker[pos] += int(bit)

    # work with gamma/epsilon as arr, easier
    gamma_arr = [None] * bits
    epsilon_arr = [None] * bits

    for pos, total_bits in enumerate(binary_tracker):
        if total_bits > threshold:
            gamma_arr[pos] = 1
            epsilon_arr[pos] = 0
        else:
            gamma_arr[pos] = 0
            epsilon_arr[pos] = 1

    # convert arr to binary string
    gamma = "".join([str(i) for i in gamma_arr])
    epsilon = "".join([str(i) for i in epsilon_arr])

    print("Part 1:", int(gamma,2) * int(epsilon, 2))
    # solution == 1082324


def part_2_helper(critera_type):
    if critera_type == "most":
        trump = "1"
        alternate = "0"
    else:
        trump = "0"
        alternate = "1"

    input_ = get_input()
    bits = len(input_[0])
    # bits is the number of bits in each binary number

    # Determine rating based on critera type
    for bit in range(bits):
        max_ = len(input_)
        # max_ is the greatest possible number of 1s or 0s for each bit
        threshold = max_ / 2
        # threshold is the number of 1s or 0s needed to win majority

        binary_tracker = 0

        for binary_number in input_:
            binary_tracker += int(binary_number[bit])

        if binary_tracker >= threshold:
            input_ = [i for i in input_ if i[bit] == trump]

        if binary_tracker < threshold:
            input_ = [i for i in input_ if i[bit] == alternate]

        if len(input_) == 1:
            break

    value_bin = "".join([str(i) for i in input_[0]])
    return int(value_bin, 2)


def part_2():
    oxygen = part_2_helper("most")
    co2 = part_2_helper("least")
    print("Part 2:", oxygen * co2)
    # solution == 1353024


def main():
    part_1()
    part_2()


if __name__ == "__main__":
    main()
