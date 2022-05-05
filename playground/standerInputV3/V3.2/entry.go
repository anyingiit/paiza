package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(GetInputWithPaizaStanderd())
}

//func ParsePaizaStanderdInput(strArray []string, firstLineArgsCount, height, width int) (args []int, datas [][]int) {
//	args = ParseStringArrayToIntArray(strArray[:1], 1, firstLineArgsCount)[0]
//	datas = ParseStringArrayToIntArray(strArray[1:], height, width)
//	return args, datas
//}

func GetInputWithPaizaStanderd() (lineHeight, lineWidth int, datas [][]int) {
	_, err := fmt.Scanf("%d %d\n", &lineHeight, &lineWidth)
	if err != nil {
		panic(fmt.Errorf("get arge for fmt.Scanf failed: %s", err.Error()))
	}
	datas = ParseStringArrayToIntArray(GetInputs(lineHeight), lineHeight, lineWidth)
	return lineHeight, lineWidth, datas
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

func GetInputs(line int) (inputs []string) {
	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < line; i++ {
		sc.Scan()
		inputs = append(inputs, sc.Text())
	}
	return inputs
}
