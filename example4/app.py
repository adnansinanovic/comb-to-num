def main():
    print(":: The bit flags")
    set = ["A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"]
    subset = ["C", "F", "H"]

    runExample(set, subset)


def runExample(set, subset):
    indices = findIndices(set, subset)
    print("subset: {0}, indices: {1}".format(subset, indices))

    encoded = encode(indices)
    print("encoded: {0}".format(encoded))

    decodedIndices = decode(encoded)
    val = values(set, decodedIndices)
    print("decodedIndices: {0}, decoded: {1}".format(decodedIndices, val))


def encode(indices):
    result = 0
    for i in range(len(indices)):
        el = indices[i]
        result += 2 ** el

    return result


def decode(encoded):
    max = 32
    result = []
    for i in range(max):
        p = 1 << i
        if ((p & encoded) == p):
            result.append(i)

    return result


def values(set, indices):
    subset = []
    for i in range(len(indices)):
        element = indices[i]
        subset.append(set[element])

    return subset


def findIndices(set, subset):
    indices = []
    for i in range(len(subset)):
        element = subset[i]
        indices.append(index(set, element))

    return indices


def index(set, element):
    for i in range(len(set)):
        el = set[i]
        if (element == el):
            return i

    raise Exception("element does not belong to the set")


main()