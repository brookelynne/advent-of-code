def part_1():
    accumulator = []
    for i in range(12):
        accumulator.append({'0': 0, '1': 0})

    with open('03_input.txt', 'r') as f:
        for line in f:
            for i, bit in enumerate(line):
                if bit == "\n":
                    continue
                accumulator[i][bit] += 1

    # gamma code is the most common bit in each position
    gamma, epsilon = get_gamma_epsilon(accumulator)
    print("Gamma: {0}, {1}".format(gamma, epsilon))
    gamma_base10 = int(gamma, 2)
    epsilon_base10 = int(epsilon, 2)
    print("Values: {0}, {1}".format(gamma_base10, epsilon_base10))
    print("Final: {0}".format(gamma_base10 * epsilon_base10))


def get_gamma_epsilon(accumulator):
    gamma = ""
    epsilon = ""
    for bit in accumulator:
        gamma_epsilon = ('0', '1') if bit['0'] > bit['1'] else ('1', '0')
        gamma = gamma + gamma_epsilon[0]
        epsilon = epsilon + gamma_epsilon[1]
    return gamma, epsilon


def part_2():
    # find most (or least) common far-left bit. Filter out numbers without that bit. Continue filtering to the right
    # until you have one number left in the list.

    accumulator = {'0': 0, '1': 0}
    numbers = []

    with open('03_input.txt', 'r') as f:
        for line in f:
            numbers.append(line)
            accumulator[line[0]] += 1

    keep = ''
    while len(numbers) > 1 and len(keep) < 10:
        keep = keep + ('0' if accumulator['0'] > accumulator['1'] else '1')
        numbers = [number for number in numbers if number[0:len(keep)] == keep]
        accumulator = {'0': 0, '1': 0}
        for num in numbers:
            accumulator[num[len(keep)]] += 1
    for num in numbers:
        print(num)


if __name__ == '__main__':
    part_2()
