BRACKET_TYPE = {
    "{": "OPEN",
    "}": "CLOSE",
    "[": "OPEN",
    "]": "CLOSE",
    "(": "OPEN",
    ")": "CLOSE",
    "<": "OPEN",
    ">": "CLOSE",
}

BRACKET_MATCH = {"{": "}", "[": "]", "(": ")", "<": ">"}

BRACKET_POINT_LOOKUP = {")": 3, "]": 57, "}": 1197, ">": 25137}
BRACKET_POINT_LOOKUP_P2 = {")": 1, "]": 2, "}": 3, ">": 4}


def get_input():
    with open("input.txt") as f:
        c = f.readlines()
    return [i.strip() for i in c]


def part_1():
    input_ = get_input()
    record = {")": 0, "]": 0, "}": 0, ">": 0}

    for line in input_:
        stack = []
        for char in line:
            # if its an open bracket, add it to the stack
            if BRACKET_TYPE[char] == "OPEN":
                stack.append(char)
            # if its a correct closing bracket, process it and continue
            elif len(stack) > 0 and BRACKET_MATCH[stack[-1]] == char:
                del stack[-1]
            # if its an illegal closing bracket, record it and break
            else:
                record[char] += 1
                break

    print(
        "Part 1:",
        (record[")"] * BRACKET_POINT_LOOKUP[")"])
        + (record["}"] * BRACKET_POINT_LOOKUP["}"])
        + (record["]"] * BRACKET_POINT_LOOKUP["]"])
        + (record[">"] * BRACKET_POINT_LOOKUP[">"]),
    )
    # solution == 436497


def part_2():
    input_ = get_input()
    completion_sequences = []
    completion_sequences_scores = []

    for line in input_:
        stack = []
        completion_sequence = []
        illegal = False

        for char in line:
            # if its an open bracket, add it to the stack
            if BRACKET_TYPE[char] == "OPEN":
                stack.append(char)
            # if its a correct closing bracket, process it and continue
            elif len(stack) > 0 and BRACKET_MATCH[stack[-1]] == char:
                del stack[-1]
            # if its an illegal closing bracket, break
            else:
                illegal = True
                break

        # if its an incomplete, then proceed!
        if not illegal and len(stack) > 0:
            while len(stack) > 0:
                completion_sequence.append(BRACKET_MATCH[stack.pop(-1)])

        if len(completion_sequence) > 0:
            completion_sequences.append(completion_sequence)

    # figure out the score of each completion sequence
    for seq in completion_sequences:
        score = 0
        for char in seq:
            score *= 5
            score += BRACKET_POINT_LOOKUP_P2[char]
        completion_sequences_scores.append(score)

    completion_sequences_scores.sort()

    # doing `len(completion_sequences_scores) // 2` gives us the middle ele since its guaranteed to be odd length (and indexs start at 0)
    print("Part 2:", completion_sequences_scores[len(completion_sequences_scores) // 2])
    # solution == 2377613374


def main():
    part_1()
    part_2()


if __name__ == "__main__":
    main()
