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

func calcFuel(f float64) float64 {
	return math.Floor(f/3.0) - 2.0
}

func calcFuelRec(f float64) int {
	value := math.Floor(f/3.0) - 2.0
	if value > 0 {
		globalsum += int(value)
		calcFuelRec(value)
	}
	return globalsum
}

var globalsum int = 0
var sum int = 0

func main() {
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
		sum += int(calcFuel(value))
		sum += int(calcFuelRec(value))

	}
	fmt.Println("Task 1 " + strconv.Itoa(sum))
	fmt.Println("Task 2 " + strconv.Itoa(globalsum))

	file.Close()
	// fmt.Println(sum)
}
