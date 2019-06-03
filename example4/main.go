package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf(":: The bit flags\n")
	set := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"}
	subset := []string{"C", "F", "H"}

	runExample(set, subset)
}

func runExample(set, subset []string) {
	indices := indices(set, subset)
	fmt.Printf("subset: %v, indices: %v\n", subset, indices)

	encoded := encode(indices)
	fmt.Printf("encoded: %v\n", encoded)

	decodedIndices := decode(encoded)
	fmt.Printf("decodedIndices: %v, decoded: %v\n", decodedIndices, values(set, decodedIndices))
}

func encode(indices []int) uint64 {
	var result uint64
	for _, v := range indices {
		result += uint64(math.Pow(float64(2), float64(v)))

		// equivalent to previous line
		// result |= (uint64(1) << uint64(v))
	}

	return result
}

func decode(encoded uint64) []int {
	max := 64 // 2^64 = math.Uint64
	result := []int{}
	for i := 0; i < max; i++ {
		p := uint64(1) << uint64(i)

		// equivalent to previous line
		// p := uint64(math.Pow(float64(2), float64(i)))

		if (p & encoded) == p {
			result = append(result, i)
		}
	}
	return result
}

func values(set []string, indices []int) []string {
	values := []string{}

	for _, v := range indices {
		values = append(values, set[v])
	}

	return values
}

func indices(set, subset []string) []int {
	indices := []int{}
	for _, v := range subset {
		indices = append(indices, index(set, v))
	}

	return indices
}

func index(set []string, element string) int {
	for i, v := range set {
		if v == element {
			return i
		}
	}

	panic("element does not belong to the set")
}
