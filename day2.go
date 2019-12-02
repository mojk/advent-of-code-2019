package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

var pos1, pos2, dest, sum int
var pos int = 0

func main() {

	file, _ := os.Open("day2.txt")
	reader := csv.NewReader(bufio.NewReader(file))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		record[1] = "12"
		record[2] = "2"

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
			fmt.Println(record[0])
		}
	}
}
