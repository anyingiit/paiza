package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputs := getInputs(true)
	fmt.Println(inputs)
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
