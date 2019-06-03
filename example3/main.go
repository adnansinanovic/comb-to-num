package main

import (
	"fmt"
	"sort"

	"github.com/gonum/stat/combin"
)

func main() {
	// set of all elements
	set := []string{"A", "B", "C", "D", "E", "F", "G", "G", "I", "J", "K"}

	// combination of elements from set
	combination := []string{"C", "K", "G"}
	runExample(set, combination)
}

func runExample(set, subset []string) {
	fmt.Println(":: Combinadics")

	// get indices for combination element
	indices := indices(set, subset)
	fmt.Printf("subset %v, indices: %v\n", subset, indices)

	// sort
	sort.Ints(indices)
	reverse(indices)
	fmt.Printf("sorted indices: %v\n", indices)

	encoded := encode(indices)
	fmt.Printf("encoded:  %v\n", encoded)

	decodedIndices := decode(encoded, len(indices), 20)

	fmt.Printf("decoded indices: %v, decoded: %v\n", decodedIndices, values(set, decodedIndices))
}

func encode(set []int) int {
	code := 0
	for i := range set {
		n := set[i] - 1
		m := len(set) - i

		if n >= m {
			code += combin.Binomial(n, m)
		}

	}
	return code
}

func decode(encodedValue, subsetLength, totalElements int) []int {
	var result []int = []int{}
	binom := combin.Binomial(totalElements, subsetLength)
	for index := totalElements; index >= 1; index-- {

		if binom <= encodedValue {
			result = append(result, index+1)
			encodedValue -= binom
			binom = binom * subsetLength / (index - subsetLength + 1)
			subsetLength--
		}
		binom = binom * (index - subsetLength) / index
	}

	return result
}

func reverse(array []int) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}

func values(set []string, indices []int) []string {
	subset := []string{}
	for _, v := range indices {
		subset = append(subset, set[v])
	}

	return subset
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
