package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const FILENAME = "input.txt"

func main() {
	var score int64

	file, err := os.Open(FILENAME)
	if err != nil {
		fmt.Println("error opening ", FILENAME)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for j := range 1000 {
		buffer, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(fmt.Sprintf("error reading line %d: %v\n", j+1, err))
		}
		if isPrefix {
			panic("buffer too short")
		}

		line := strings.Split(string(buffer), " ")
		if checkSafety(line, j) {
			score++
			continue
		} else {
			for k := 0; k < len(line); k++ {
				originalElement := line[k]
				line = slices.Delete(line, k, k+1)
				if checkSafety(line, j) {
					score++
					break
				}
				line = slices.Insert(line, k, originalElement)
			}
		}
	}

	fmt.Println(score)
}

func checkSafety(line []string, j int) bool {
	var ascending bool
	var x1, x2 int64
	var err error

	fmt.Println(line)
	x1, err = strconv.ParseInt(line[0], 10, 64)
	if err != nil {
		panic(fmt.Sprintf("error in conversion on line %d", j))
	}
	for i := 1; i < len(line); i++ {
		x2, err = strconv.ParseInt(line[i], 10, 64)
		if err != nil {
			panic(fmt.Sprintf("error in conversion on line %d", i))
		}
		if i == 1 {
			ascending = (x1 < x2)
		}
		if x1 == x2 || math.Abs(float64(int64(x1-x2))) > 3 {
			return false
		} else if ascending && x1 > x2 {
			return false
		} else if !ascending && x1 < x2 {
			return false
		} else if i == len(line)-1 {
			return true
		}
		x1 = x2
	}
	return false
}
