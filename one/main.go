package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"testing"
)

const FILENAME = "input.txt"
const LINE_LENGTH = 13

func main() {
	list1 := make([]uint32, 1000)
	list2 := make([]uint32, 1000)
	var value, diff uint64
	var sub, mask uint32

	file, err := os.Open(FILENAME)
	if err != nil {
		fmt.Println("error opening ", FILENAME)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for i := range 1000 {
		buffer, isPrefix, err := reader.ReadLine()
		if err != nil {
			panic(fmt.Sprintf("error reading line %d: %v\n", i+1, err))
		}
		if isPrefix {
			panic("buffer too short")
		}

		value, err = strconv.ParseUint(string(buffer[0:5]), 10, 32)
		if err != nil {
			panic(fmt.Sprintf("error in conversion on line %d", i))
		}
		list1[i] = uint32(value)
		
		value, err = strconv.ParseUint(string(buffer[8:13]), 10, 32)
		if err != nil {
			panic(fmt.Sprintf("error in conversion on line %d", i))
		}
		list2[i] = uint32(value)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	for i := range 1000 {
		sub = list1[i] - list2[i]
		mask = uint32(int32(sub) >> 31)
		diff = diff + uint64((sub^mask)-mask)
	}

	fmt.Println(diff)
}


// Below code is for benchmarking ONLY. 
// Benchmarks run on Apple M2 16 GB. 
// Observations: 

// Least performant (by 0.01 ns/op at most)
// 0 B/op 0 allocs/op
func absDiffUint32_conditional(a, b uint32) uint32 {
	if a > b {
		return a - b
	}
	return b - a
}

// Most performant
// 0 B/op 0 allocs/op
func absDiffUint32_bitManipulation(a, b uint32) uint32 {
	sub := a - b
	mask := uint32(int32(sub) >> 31)
	return (sub ^ mask) - mask
}

// Less performant than bitManipulation (by 0.001 ns/op at most)
// 0 B/op 0 allocs/op
func absDiffUint32_builtin(a, b uint32) uint32 {
	return uint32(math.Abs(float64(a-b)))
}

func BenchmarkAbsDiffConditional(b *testing.B) {
	a := uint32(12345)
	c := uint32(54321)
	for i := 0; i < b.N; i++ {
		_ = absDiffUint32_conditional(a+uint32(i), c-uint32(i))
	}
}

func BenchmarkAbsDiffBitManipulation(b *testing.B) {
	a := uint32(12345)
	c := uint32(54321)
	for i := 0; i < b.N; i++ {
		_ = absDiffUint32_bitManipulation(a+uint32(i), c-uint32(i))
	}
}

func BenchmarkAbsDiffBuiltIn(b *testing.B) {
	a := uint32(12345)
	c := uint32(54321)
	for i := 0; i < b.N; i++ {
		_ = absDiffUint32_builtin(a+uint32(i), c-uint32(i))
	}
}
