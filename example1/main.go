package main

import "fmt"

var primes []uint64
var primeIndices map[uint64]uint64

func main() {
	// prepare prime numbers
	primes = sieveOfEratosthenes(200)
	primeIndices = map[uint64]uint64{}
	for i, v := range primes {
		primeIndices[v] = uint64(i)
	}

	// set
	set := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"}

	// subset indices
	example1 := []uint64{1, 3, 4}
	example2 := []uint64{6, 3, 9, 10}
	example3 := []uint64{6, 3, 9, 10, 7, 1}

	runExample(set, example1, "Example 1")
	runExample(set, example2, "Example 2")
	runExample(set, example3, "Example 3")
}

func runExample(set []string, subsetIndices []uint64, title string) {
	fmt.Printf(":: Primes\n")
	fmt.Printf("Running: %v :: %v\n", title, subsetIndices)
	encoded := encode(subsetIndices)
	decoded := decode(encoded)

	fmt.Printf("Subset %v is represented with number: %v\n", debug(set, subsetIndices), encoded)
	fmt.Printf("Number: %v represents: %v -> %v\n\n", encoded, decoded, debug(set, decoded))
}

func decode(n uint64) []uint64 {
	pfs := []uint64{}
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := uint64(3); i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}
	// return pfs

	indices := []uint64{}
	for _, v := range pfs {
		indices = append(indices, primeIndices[v])
	}

	return indices
}

func encode(indices []uint64) uint64 {
	primes := sieveOfEratosthenes(200)

	var product uint64
	product = 1

	for _, v := range indices {
		product *= uint64(primes[v])
	}

	return product
}

// sieveOfEratosthenes return list of primes less than N
func sieveOfEratosthenes(N int) (p []uint64) {
	b := make([]bool, N)
	for i := 2; i < N; i++ {
		if b[i] == true {
			continue
		}
		p = append(p, uint64(i))
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}

	return p
}

func debug(set []string, indices []uint64) string {
	subset := []string{}
	for _, v := range indices {
		subset = append(subset, set[v])
	}

	return fmt.Sprintf("%v", subset)
}
