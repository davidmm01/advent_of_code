def get_input():
    with open("input.txt") as f:
        contents = f.readlines()
    output = []
    for i in contents:
        i = i.strip()
        line = []
        for j in i:
            line.append(int(j.strip()))
        output.append(line)
    return output

def part_1():
    input_ = get_input()
    print("Part 1:")
    # solution == 


def part_2():
    print("Part 2:")
    # solution ==

def main():
    part_1()
    part_2()


if __name__ == "__main__":
    main()
