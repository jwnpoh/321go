package main

import (
	"fmt"
	"os"
	"strconv"
)

const usage = `
A simple countdown timer. 
Usage: 
321go hh mm ss
`

func main() {
	input := getInput()
	output := parseInput(input)

	var hh, mm, ss int

	switch len(output) {
	case 1:
		ss = output[0]
	case 2:
		mm, ss = output[0], output[1]
	case 3:
		hh, mm, ss = output[0], output[1], output[2]
	}

	fmt.Println(hh, mm, ss)
}

func getInput() []string {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Print(usage)
		os.Exit(1)
	}
	return os.Args[1:]
}

func parseInput(input []string) []int {
	var o []int
	for _, j := range input {
		n, _ := strconv.Atoi(j)
		o = append(o, n)
	}
	return o
}

func getDuration() {

}
