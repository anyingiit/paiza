package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// V3.31 fixed I know all err
func main() {
	datas, height, width := GetInputWithPaizaStanderd()
	fmt.Println(height, width, datas)
}

func GetInputWithPaizaStanderd() (datas [][]int, lineHeight, lineWidth int) {
	lineHeight, lineWidth = GetInputWithPaizaStanderedArgs()
	datas = GetInputWithPaizaStanderedDatas(lineHeight, lineWidth)
	return datas, lineHeight, lineWidth
}

func GetInputWithPaizaStanderedDatas(height, width int) (datas [][]int) {
	sc := bufio.NewScanner(os.Stdin)

	datas = make([][]int, height)
	for i := 0; i < height; i++ {
		datas[i] = make([]int, width)
		sc.Scan()
		inputStr := sc.Text()
		curStr := sc.Text()

		for j := 0; j < width; j++ {
			_, err := fmt.Sscanf(curStr, "%d ", &datas[i][j])
			if err != nil {
				panic(fmt.Errorf(
					"failed scan string %s, now string status is %s, i=%d j=%d, width=%d height=%d, originError: %s",
					inputStr, curStr, i, j, width, height, err.Error()))
			}
			nextStartIndex := len(strconv.Itoa(datas[i][j]))
			if j != width-1 {
				nextStartIndex++
			}
			curStr = curStr[nextStartIndex:]
		}
	}
	return datas
}

func GetInputWithPaizaStanderedArgs() (lineHeight, lineWidth int) {
	_, err := fmt.Scanf("%d %d\n", &lineHeight, &lineWidth)
	if err != nil {
		panic(fmt.Errorf("get arge for fmt.Scanf failed: %s", err.Error()))
	}
	return lineHeight, lineWidth
}
