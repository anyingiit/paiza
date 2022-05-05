package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputs := getInputs(false)
	ints := parseInputsToInts(inputs)
	fmt.Println(ints[0] * 12)
}

func parseInputsToInts(inputs []string) (ints []int) {
	for _, s := range inputs {
		atoi, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Errorf("str to int failed: %s", err.Error()))
		}
		ints = append(ints, atoi)
	}
	return ints
}

func getInputs(stopInputWithWordStop bool) (inputs []string) {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		if stopInputWithWordStop && sc.Text() == "stop" {
			break
		}
		inputs = append(inputs, sc.Text())
	}
	return inputs
}
