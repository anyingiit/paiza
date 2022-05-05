package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputs := getInputs(false)
	n, _ := strconv.Atoi(inputs[0])
	m, _ := strconv.Atoi(inputs[1])
	fmt.Println(n * m)
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
