package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var x, y int
	inputs := getInputs(false)
	fmt.Sscanf(inputs[0], "%d %d", &x, &y)

	fmt.Println(x - y)
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
