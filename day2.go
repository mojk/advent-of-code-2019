package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func inputFromFile() []string {

	file, _ := os.Open("day2.txt")
	defer file.Close()
	reader := csv.NewReader(bufio.NewReader(file))
	input, _ := reader.Read()
	return input
}

var pos1, pos2, dest, sum int
var record []string

func compute(input string, input2 string, slice []string) string {

	var pos int = 0
	record[1] = input
	record[2] = input2

	for cond := true; cond; cond = record[pos] != "99" {

		switch record[pos] {
		case "1":
			pos1, _ = strconv.Atoi(record[pos+1])
			pos2, _ = strconv.Atoi(record[pos+2])

			newval, _ := strconv.Atoi(record[pos1])
			newval2, _ := strconv.Atoi(record[pos2])
			sum = newval + newval2

			dest, _ = strconv.Atoi(record[pos+3])

			record[dest] = strconv.Itoa(sum)
			pos += 4

		case "2":
			pos1, _ = strconv.Atoi(record[pos+1])
			pos2, _ = strconv.Atoi(record[pos+2])

			newval, _ := strconv.Atoi(record[pos1])
			newval2, _ := strconv.Atoi(record[pos2])
			sum = newval * newval2

			dest, _ = strconv.Atoi(record[pos+3])
			record[dest] = strconv.Itoa(sum)
			pos += 4
		default:
		}

	}
	return record[0]
}

func main() {
	/* Task 1*/
	record = inputFromFile()
	output := compute("12", "2", record)
	fmt.Println("First task " + output)

	/* Task 2 */
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			record = inputFromFile()
			output := compute(strconv.Itoa(i), strconv.Itoa(j), record)
			if output == "19690720" {
				fmt.Println(i)
				fmt.Println(j)
				break
			}

		}
	}
}
