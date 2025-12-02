
def part_1():
    previous_depth = None
    deeper = 0
    with open('01_input.txt', 'r') as f:
        for depth in f:
            depth = int(depth)
            if previous_depth is not None and depth > previous_depth:
                deeper += 1
            previous_depth = depth

    f.close()

    print("measurements that are larger than the previous one: {0}".format(deeper))


def part_2():
    previous_summation = None
    summation = None
    deeper = 0

    before_previous = None
    previous = None
    with open('01_input.txt', 'r') as f:
        for current in f:
            current = int(current)

            if before_previous is not None and previous is not None:
                summation = before_previous + previous + current

            if previous_summation is not None and summation is not None and summation > previous_summation:
                deeper += 1

            before_previous = previous
            previous = current
            previous_summation = summation

    f.close()

    print("measurements that are larger than the previous one: {0}".format(deeper))


if __name__ == "__main__":
    part_2()
