def part_1():
    position = [0, 0]
    with open('02_input.txt', 'r') as f:
        for command in f:
            parts = command.split()
            val = int(parts[1])
            if parts[0] == 'forward':
                position[1] += val
            elif parts[0] == 'up':
                position[0] -= val
            elif parts[0] == 'down':
                position[0] += val
            else:
                print("not forward, up, or down??")
    print("Position: {0}, {1}".format(position[0], position[1]))
    print("Final: {0}".format(position[0] * position[1]))


def part_2():
    position = [0, 0]
    aim = 0
    with open('02_input.txt', 'r') as f:
        for command in f:
            parts = command.split()
            val = int(parts[1])
            if parts[0] == 'forward':
                position[1] += val
                position[0] += aim * val
            elif parts[0] == 'up':
                aim -= val
            elif parts[0] == 'down':
                aim += val
            else:
                print("not forward, up, or down??")
    print("Position: {0}, {1}; Aim: {2}".format(position[0], position[1], aim))
    print("Final: {0}".format(position[0] * position[1]))


if __name__ == '__main__':
    part_2()
