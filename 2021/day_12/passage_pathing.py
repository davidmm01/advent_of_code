START = "start"
END = "end"


def get_data_from_file(filename):
    # tunnels will a dict, where the key is the source tunnel label, and the
    # value is all tunnels that begin from the source tunnel
    tunnels = {}
    with open(filename) as f:
        lines = f.readlines()
    for line in lines:
        line = line.strip()
        src, dest = line.split("-")
        if tunnels.get(src) == None:
             tunnels[src] = []
        tunnels[src].append(dest)
    return tunnels


def traverse_tunnels(tunnels):
    for path in tunnels[START]:
        

def part_1():
    # sample 1
    traverse_tunnels(get_data_from_file("sample_1.txt"))

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
