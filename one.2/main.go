package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const FILENAME = "input.txt"
const LINE_LENGTH = 13

func main() {
	map1 := make(map[uint64]uint64, 1000)
	map2 := make(map[uint64]uint64, 1000)
	var value, score uint64

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
		map1[value]++

		value, err = strconv.ParseUint(string(buffer[8:13]), 10, 32)
		if err != nil {
			panic(fmt.Sprintf("error in conversion on line %d", i))
		}
		map2[value]++
	}

	for k, v := range map1 {
		score += k * v * map2[k]
	}

	fmt.Println(score)
}
