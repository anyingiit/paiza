package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputs := GetInputs(true)
	input, width, datas := ParsePaizaStanderdInput(inputs)
	fmt.Println(input, width, datas)
}

//func ParsePaizaStanderdInput(strArray []string, firstLineArgsCount, height, width int) (args []int, datas [][]int) {
//	args = ParseStringArrayToIntArray(strArray[:1], 1, firstLineArgsCount)[0]
//	datas = ParseStringArrayToIntArray(strArray[1:], height, width)
//	return args, datas
//}

func ParsePaizaStanderdInput(strArray []string) (lineHeight, lineWidth int, datas [][]int) {
	args := ParseStringArrayToIntArray(strArray[:1], 1, 2)[0]
	datas = ParseStringArrayToIntArray(strArray[1:], args[0], args[1])
	return args[0], args[1], datas
}

func ParseStringArrayToIntArray(strArray []string, height, width int) (result [][]int) {
	result = make([][]int, height)
	for i := 0; i < height; i++ {
		result[i] = make([]int, width)
		str := strArray[i]
		for j := 0; j < width; j++ {
			n, err := fmt.Sscanf(str, "%d ", &result[i][j])
			if err != nil {
				panic(fmt.Errorf(
					"failed scan string %s, now string status is %s, i=%d j=%d, width=%d height=%d, originError: %s",
					strArray[i], str, i, j, width, height, err.Error()))
			}
			nextStartIndex := n
			if j != width-1 {
				nextStartIndex++
			}
			str = str[nextStartIndex:]
		}
	}
	return result

}

func GetInputs(stopInputWithWordStop bool) (inputs []string) {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		if stopInputWithWordStop && sc.Text() == "stop" {
			break
		}
		inputs = append(inputs, sc.Text())
	}
	return inputs
}
