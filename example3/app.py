def main():
    # set of all elements
    set = ["A", "B", "C", "D", "E", "F", "G", "G", "I", "J", "K"]

    # combination of elements from set
    combination = ["C", "K", "G"]
    runExample(set, combination)


def runExample(set, subset):
    print(":: Combinadics")

    # get indices for combination element
    indices = findIndices(set, subset)
    print("subset {0}, indices: {1}".format(subset, indices))

    # sort
    indices.sort()
    print("sorted indices: {0}".format(indices))

    encoded = encode(indices, len(set))
    print("encoded: {0}".format(encoded))

    decodedIndices = decode(encoded, len(indices), 20)
    val = values(set, decodedIndices)
    print("decoded indices: {0}, decoded: {1}".format(decodedIndices, val))


def encode(set, totalElements):
    code = 0
    for i in range(len(set)):
        n = totalElements - set[i]
        # decreasing k by one (..., 3, 2, 1) for the lower part of binomial
        # coefficient
        m = len(set) - i
        # calculate the binomial coefficient and add.
        code += binomial(n, m)
        
    return code


def decode(encodedValue, subsetLength, totalElements):
    result = []
    binom = binomial(totalElements, subsetLength)

    index = totalElements
    while index >= 1:
        if binom <= encodedValue:
            result.append(index + 1)
            encodedValue -= binom
            binom = binom * subsetLength // (index - subsetLength + 1)
            subsetLength -= 1
        binom = binom * (index - subsetLength) // index
        index -= 1

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


def binomial(n, k):
    coeff = 1

    if k < 0:
        return 0

    if k > n:
        return 0

    for i in range(k):
        coeff = coeff * (n - i) // (i + 1)

    return coeff


main()
