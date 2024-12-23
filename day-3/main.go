package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"zacharylyh/aoc2024"
)

func main() {
	file := aoc2024.ExtractTxtFile()
	defer file.Close()

	scanner := bufio.NewScanner(file)

	partOneTotalSum := 0
	partTwoTotalSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		rows := strings.Fields(line)

		str := ""
		for _,r := range rows {
			str += r
		}
		res, _ := partOne(str, false)
		partOneTotalSum += res
		partTwoTotalSum += partTwo(str)
	}
	fmt.Println("Part one: ", partOneTotalSum)
	fmt.Println("Part two: ", partTwoTotalSum)
}

func partTwo(str string) int {
	sum := 0
	for i:=0; i<len(str); i++ {
		if i == 0 {
			res, nextStart := partOne(str, true)
			i = nextStart
			sum += res
			continue
		}
		if i+4 < len(str) && str[i] == 'd' && str[i+1] == 'o' && str[i+2] == '(' && str[i+3] == ')' {
			res, nextStart := partOne(str[i+4:], true)
			i = i+nextStart
			sum += res
		}
	} 
	return sum
}

func partOne(str string, isPartTwo bool) (int, int) {
	sum := 0
	for i:=0; i<len(str); i++ {
		if isPartTwo {
			if i+7 < len(str) && str[i:i+7] == "don't()" {
				return sum, i+7
			}
		}
		if str[i] == 'm' {
			potential := extractPotentialMul(str[i:])
			mul := checkIfMulUsable(potential)
			if mul != "" {
				sum += evaluateMulReturnNum(mul)
				i += len(mul)
			} 
		}
	}
	return sum, len(str)
}

func evaluateMulReturnNum(mul string) int {
	num1 := 0
	num2 := 0
	commaSeen := false
	for _,m := range mul {
		if m == 44 {
			commaSeen = true
			continue
		}
		num, err := strconv.Atoi(string(m))
		aoc2024.HandleGenericErr(err)
		if !commaSeen {
			num1 = (num1*10)+num
		} else {
			num2 = (num2*10)+num
		}
	}
	return num1*num2
}

/*
happy path: a,b)
don't allow anything other than numbers or ","
if output is a string, guaranteed to look like the happy path and no malformed mul
*/
func checkIfMulUsable(mul string) string { 
	potentialInner := ""
	commaFreq := 0
	closeParenFreq := 0
	for _,m := range mul {
		if (m >= 48 && m <= 57) || m == 44 || m == 41{
			if m == 44 {
				if commaFreq > 1 {
					return ""
				} 
				commaFreq++
			}
			if m == 41 {
				if closeParenFreq > 1 {
					return ""
				} 
				closeParenFreq++
			}
			potentialInner += string(m)
		} else {
			return ""
		}
	}
	return potentialInner
}

func extractPotentialMul(str string) string {
	if len(str) < 9 { //mul(a,b) 8 characters
		return ""
	}
	if str[0] == 'm' && str[1] == 'u' && str[2] == 'l' && str[3] == '(' {
		i := 4
		endedWithClosingParenthesis := false
		for i=4; i<len(str); i++ {
			if str[i] == 41 { // )
				endedWithClosingParenthesis = true //if we don't even end with a closing paren then no point returning
				break
			}
		}
		if endedWithClosingParenthesis {
			return str[4:i]
		} else {
			return ""
		}
	}
	return ""
}