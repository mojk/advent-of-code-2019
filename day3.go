package main

import (
	"bufio"
	"fmt"
	"os"
)

func inputFromFile() ([]string, []string) {

	var w1, w2 []string
	file, _ := os.Open("day3.txt")
	defer file.Close()
	reader := bufio.NewReader(file)

	w1 = reader.ReadLine()
	w2 = reader.ReadLine()

	return w1, w2
}

var wire1, wire2 []string

func main() {
	wire1, wire2 := inputFromFile()
	fmt.Println(wire1)
	fmt.Println(wire2)
}
