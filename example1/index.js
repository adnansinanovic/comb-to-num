let primes = []
let primeIndices = []
function main() {
    primes = sieveOfEratosthenes(200);
    primeIndices = {};
    for (let i = 0; i < primes.length; i++) {
        const element = primes[i];
        primeIndices[element] = i
    }
    
    // set
    const set = ["A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"];
    
    // subset indices
    const example1 = [1, 3, 4];
    const example2 = [6, 3, 9, 10];
    const example3 = [6, 3, 9, 10, 7, 1];
    
    runExample(set, example1, "Example 1")
    runExample(set, example2, "Example 2")
    runExample(set, example3, "Example 3")
}


function runExample(set, subsetIndices, title) {
    console.log(":: Primes")
    console.log(`Running: ${title} :: [${subsetIndices}]`)
    const encoded = encode(subsetIndices)
    const decoded = decode(encoded)

    console.log(`Subset [${values(set, subsetIndices)}] is represented with number: ${encoded}}`);
    console.log(`Number: ${encoded} represents: [${decoded}] -> [${values(set, decoded)}]\n`);
}

function decode(n) {
    var pfs = [];
    // Get the number of 2s that divide n
    while (n % 2 == 0) {
        pfs.push(2)
        n = Math.floor(n / 2);
    }

    // n must be odd at this point. so we can skip one element
    // (note i = i + 2)
    for (let i = 3; i * i <= n; i = i + 2) {
        while (n % i == 0) {
            pfs.push(i)
            n = Math.floor(n / i)
        }
    }

    // This condition is to handle the case when n is a prime number
    // greater than 2
    if (n > 2) {
        pfs.push(n)
    }
    // return pfs

    const indices = [];
    for (let i = 0; i < pfs.length; i++) {
        const element = pfs[i];
        indices.push(primeIndices[element])

    }

    return indices
}

function encode(indices) {
    const primes = sieveOfEratosthenes(200)

    var product = 1
    for (let i = 0; i < indices.length; i++) {
        const element = primes[indices[i]];
        product *= element;
    }


    return product
}

// sieveOfEratosthenes return list of primes less than N
function sieveOfEratosthenes(N) {
    const p = [];
    const b = new Array(N);

    for (let i = 2; i < b.length; i++) {
        if (b[i] == true) {
            continue;
        }

        p.push(i);
        for (let k = i * i; k < N; k += i) {
            b[k] = true
        }
    }

    

    return p
}

function values(set, indices) {
    const subset = [];
    for (let i = 0; i < indices.length; i++) {
        const element = indices[i];
        subset.push(set[element])
    }

    return subset
}


main();