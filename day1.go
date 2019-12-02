package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func calc(f float64) float64 {
	return math.Floor(f/3.0) - 2.0
}

func main() {
	var sum int = 0
	file, err := os.Open("day1.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {

		if scanner.Text() == "" {
			break
		}

		value, err := strconv.ParseFloat(scanner.Text(), 64)
		check(err)
		sum += int(calc(value))
	}

	file.Close()
	fmt.Println(sum)
}
