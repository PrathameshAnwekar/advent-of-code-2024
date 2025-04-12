package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

const FILENAME = "input.txt"

func main() {
	var x1, x2, score int64
	var flag bool

	file, err := os.Open(FILENAME)
	if err != nil {
		fmt.Println("error opening ", FILENAME)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for j := range 1000 {
		buffer, isPrefix, err := reader.ReadLine()
		fmt.Println(string(buffer))
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
		x1, _ = strconv.ParseInt(line[0], 10, 64)
		for i := 1; i < len(line); i++ {
			x2, err = strconv.ParseInt(line[i], 10, 64)
			if err != nil {
				panic(fmt.Sprintf("error in conversion on line %d", i))
			}
			if i == 1 {
				flag = (x1 < x2)
			}

			if x1 == x2 || math.Abs(float64(int64(x1-x2))) > 3 {
				fmt.Println("broke at abs", x1, x2, math.Abs(float64(x1-x2)), math.Abs(float64(x1-x2)) > 3)
				break
			} else if flag && x1 > x2 {
				fmt.Println("broke at flag")
				break
			} else if !flag && x1 < x2 {
				fmt.Println("broke at !flag")
				break
			} else if i == len(line) - 1 {
				score++
				break
			}
			x1 = x2
		}
	}

	fmt.Println(score)
}
