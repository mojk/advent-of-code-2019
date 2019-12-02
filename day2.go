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

		// for _, element := range record {
		// 	fmt.Println("Position " + strconv.Itoa(pos))
		// 	fmt.Println("Value " + element)
		for cond := true; cond; cond = record[pos] != "99" {

			switch record[pos] {
			case "1":

				fmt.Println("--- Opcode 1 found --- ")
				pos1, _ = strconv.Atoi(record[pos+1]) //val at ind
				fmt.Println("Position of value " + strconv.Itoa(pos1) + " value: " + record[pos1])

				pos2, _ = strconv.Atoi(record[pos+2]) //val at ind
				fmt.Println("Position of value " + strconv.Itoa(pos2) + " value: " + record[pos2])

				newval, _ := strconv.Atoi(record[pos1])
				newval2, _ := strconv.Atoi(record[pos2])
				sum = newval + newval2

				dest, _ = strconv.Atoi(record[pos+3]) //i to store
				fmt.Println("Adding them together and storing them at index " + strconv.Itoa(dest) + " value " + strconv.Itoa(sum))

				record[dest] = strconv.Itoa(sum)
				pos += 4
				fmt.Println(record)
				break

			case "2":
				fmt.Println("--- Opcode 2 found ---")
				fmt.Println("Parsing values located at indexes:")
				fmt.Println(pos + 1)
				fmt.Println(pos + 2)
				fmt.Println(pos + 3)

				pos1, _ = strconv.Atoi(record[pos+1]) //val at ind
				fmt.Println("Position of value " + strconv.Itoa(pos1) + " value: " + record[pos1])

				pos2, _ = strconv.Atoi(record[pos+2]) //val at ind
				fmt.Println("Position of value " + strconv.Itoa(pos2) + " value: " + record[pos2])

				newval, _ := strconv.Atoi(record[pos1])
				newval2, _ := strconv.Atoi(record[pos2])
				sum = newval * newval2

				dest, _ = strconv.Atoi(record[pos+3]) //i to store
				fmt.Println("Multiplying them together and storing them at index " + strconv.Itoa(dest) + " value " + strconv.Itoa(sum))

				record[dest] = strconv.Itoa(sum)
				pos += 4
				fmt.Println(record)
				break
			default:
			}
		}
	}
}
