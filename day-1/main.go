package main

import (
	"bufio"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"zacharylyh/aoc2024"
)

func main() {

	file := aoc2024.ExtractTxtFile()
	defer file.Close()

	var col1Arr []int
	var col2Arr []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Fields(line)
		if len(columns) != 2 {
			fmt.Println("Skipping malformed line:", line)
			continue
		}

		col1, err1 := strconv.Atoi(columns[0])
		col2, err2 := strconv.Atoi(columns[1])
		aoc2024.HandleGenericErr(err1)
		aoc2024.HandleGenericErr(err2)

		col1Arr = append(col1Arr, col1)
		col2Arr = append(col2Arr, col2)
	}

	aoc2024.HandleGenericErr(scanner.Err())

	sort.Ints(col1Arr)
	sort.Ints(col2Arr)
	diffSum := 0
	for i:=0; i<len(col1Arr); i++ {
		diffSum += int(math.Abs(float64(col1Arr[i]-col2Arr[i])))
	}
	fmt.Println("Part 1: ", diffSum)

	freqOfNumber := make(map[int]int)
	sum := 0

	for _,lVal := range col1Arr {
		if _,exists := freqOfNumber[lVal]; !exists {
			rFreq := 0
			for _,rVal := range col2Arr {
				if rVal == lVal {
					rFreq++
				}
				if rVal > lVal {
					break
				}
			}
			freqOfNumber[lVal] = rFreq
		}
		sum += lVal * freqOfNumber[lVal]
	}

	fmt.Println("Part 2: ", sum)
}

