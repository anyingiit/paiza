package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	height, width, datas := GetInputWithPaizaStanderd()
	fmt.Println(height, width, datas)
}

func GetInputWithPaizaStanderd() (lineHeight, lineWidth int, datas [][]int) {
	lineHeight, lineWidth = GetInputWithPaizaStanderedArgs()
	datas = GetInputWithPaizaStanderedDatas(lineHeight, lineWidth)
	return lineHeight, lineWidth, datas
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
			n, err := fmt.Sscanf(curStr, "%d ", &datas[i][j])
			if err != nil {
				panic(fmt.Errorf(
					"failed scan string %s, now string status is %s, i=%d j=%d, width=%d height=%d, originError: %s",
					inputStr, curStr, i, j, width, height, err.Error()))
			}
			nextStartIndex := n
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
