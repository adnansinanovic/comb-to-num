package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf(":: Odometer effect\n")
	setA := []string{"John", "Jack", "Joe"}
	setB := []string{"Water", "Fire"}
	setC := []string{"Brazil", "UK", "Russia", "Japan", "Morocco", "France"}
	setD := []string{"Blue", "Red", "Yellow", "White"}

	lengths := []int{len(setA), len(setB), len(setC), len(setD)}
	maxCombinations := calcMax(lengths)

	fmt.Printf("%3v => %8v =>  [%28v] => %4v\n", "Num", "Decoded", "Elements", "Encoded")
	fmt.Println(strings.Repeat("-", 70))

	for index := uint64(0); index < maxCombinations; index++ {
		n, err := decode(uint64(index), lengths)
		if err != nil {
			println(err.Error())
			break
		}

		elements := []uint64{n[0], n[1], n[2], n[3]}
		encoded, err := encode(elements, lengths)

		if err != nil {
			println(err.Error())
			break
		}

		fmt.Printf("%3v => %v => [%4v, %5v, %7v, %6v] => %4v\n", index, n, setA[n[0]], setB[n[1]], setC[n[2]], setD[n[3]], encoded)
	}
}

func decode(code uint64, lengths []int) ([]uint64, error) {
	max := calcMax(lengths)
	if code >= max {
		return []uint64{}, fmt.Errorf("combination index=%v is larger than max possible index=%v", code, max)
	}

	weights := calculateWeights(lengths)
	elements := make([]uint64, len(lengths))
	sums := make([]uint64, len(lengths))

	ln := len(weights) - 1
	for index := ln; index >= 0; index-- {
		if index == ln {
			elements[index] = code / weights[index]
			sums[index] = elements[index] * weights[index]
			continue
		}

		elements[index] = (code - sums[index+1]) / weights[index]
		sums[index] = sums[index+1] + elements[index]*weights[index]
	}

	return elements, nil
}

func encode(elements []uint64, lengths []int) (uint64, error) {
	if len(elements) != len(lengths) {
		return 0, fmt.Errorf("number of element (%v) is different than number of lengths (%v)", len(elements), len(lengths))
	}

	weights := calculateWeights(lengths)

	code := uint64(0)
	for i := range elements {
		code += (uint64(elements[i]) * weights[i])
	}

	return code, nil
}

func calcMax(lengths []int) uint64 {
	max := uint64(1)
	for i := range lengths {
		max *= uint64(lengths[i])
	}

	return max
}

func calculateWeights(lengths []int) []uint64 {
	weights := make([]uint64, len(lengths))
	for i := range lengths {
		if i == 0 {
			weights[0] = 1
			continue
		}

		weights[i] = weights[i-1] * uint64(lengths[i-1])
	}

	return weights
}
