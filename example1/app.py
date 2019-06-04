

primes = []
primeIndices = {}


def main(): 
    # prepare prime numbers
    primes = sieveOfEratosthenes(200)
    for i in range(len(primes)):
        primeIndices[primes[i]] = i

    # // set
    set = ["A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"]

    # // subset indices
    example1 = [1, 3, 4]
    example2 = [6, 3, 9, 10]
    example3 = [6, 3, 9, 10, 7, 1]

    runExample(set, example1, "Example 1")
    runExample(set, example2, "Example 2")
    runExample(set, example3, "Example 3")


def runExample(set, subsetIndices, title):
    print(":: Primes")
    print("Running {0} :: {1}".format(title, subsetIndices))
    encoded = encode(subsetIndices)
    decoded = decode(encoded)

    num = values(set, subsetIndices)
    print("Subset {0} is represented with number: {1}" .format(num, encoded))
    
    num = values(set, decoded)
    print("Number: {0} represents: {1} -> {2}\n".format(encoded, decoded, num))


def decode(n):
    pfs = []

    while n % 2 == 0:
        pfs.append(2)
        n = n // 2

    i = 3
    while i * i <= n:
        while n % i == 0:
            pfs.append(i)
            n = n // i
        i += 2

    if n > 2:
        pfs.append(n)

    indices = []
    for v in pfs:
        indices.append(primeIndices[v])

    return indices


def encode(indices):
    # primes = sieveOfEratosthenes(200)
    product = 1
    for v in indices:
        product *= primes[v]

    return product


def sieveOfEratosthenes(N):
    b = [None] * N
    p = []
    for i in range(2, N):
        if b[i] is True:
            continue
        p.append(i)

        for k in range(i*i, N, i):
            b[k] = True

    return p


def values(set, indices):
    subset = []
    for v in indices:
        subset.append(set[v])

    return subset


main()
