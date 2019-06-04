def main():
    print(":: Odometer effect")
    setA = ["John", "Jack", "Joe"]
    setB = ["Water", "Fire"]
    setC = ["Brazil", "UK", "Russia", "Japan", "Morocco", "France"]
    setD = ["Blue", "Red", "Yellow", "White"]

    lengths = [len(setA), len(setB), len(setC), len(setD)]
    maxCombinations = calcMax(lengths)

    print("Num => Decoded => {0} => {1}".format("Elements", "Encoded"))
    for index in range(maxCombinations):
        n = decode(index, lengths)
        elements = [n[0], n[1], n[2], n[3]]
        encoded = encode(elements, lengths)

        el1 = setA[n[0]]
        el2 = setB[n[1]]
        el3 = setC[n[2]]
        el4 = setD[n[3]]
        print("{0} => {1} => [{2:>4}, {3:>5}, {4:>7}, {5:>6}] => {6}".format(index, n, el1, el2, el3, el4, encoded))


def decode(code, lengths):
    weights = calculateWeights(lengths)
    elements = [None] * len(lengths)
    sums = [None] * len(lengths)

    ln = len(weights) - 1

    index = ln
    while index >= 0:
        if index == ln:
            elements[index] = code // weights[index]
            sums[index] = elements[index] * weights[index]
            index = index - 1
            continue
        elements[index] = (code - sums[index+1]) // weights[index]
        sums[index] = sums[index+1] + elements[index]*weights[index]
        index = index - 1

    return elements


def encode(elements, lengths):
    weights = calculateWeights(lengths)

    code = 0
    for i in range(0, len(elements)):
        code += (elements[i] * weights[i])

    return code


def calcMax(lengths):
    max = 1
    for v in lengths:
        max *= v

    return max


def calculateWeights(lengths):
    weights = [None] * len(lengths)
    for i in range(0, len(lengths)):
        if i == 0:
            weights[0] = 1
            continue

        weights[i] = weights[i-1] * lengths[i-1]

    return weights


main()

