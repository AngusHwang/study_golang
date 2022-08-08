package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	testArr := generateArr("243--2")
	for i := 0; i < len(testArr); i++ {
		fmt.Println(testArr[i])
	}
}

func evaluate(expression string) float64 {
	var (
		levelInfo []string
		resultStr string
		strA      string
		strB      string
		strC      string
	)
	resultStr = expression + " "

	for strings.Contains(resultStr, " ") || strings.Contains(resultStr, "(") || strings.Contains(resultStr, ")") || strings.Contains(resultStr, "sin") || strings.Contains(resultStr, "cos") || strings.Contains(resultStr, "tan") {
		levelInfo = checkHighestLevel(resultStr)
		strA = resultStr
		strB = calculateStr(levelInfo[0])
		tmp1, _ := strconv.ParseInt(levelInfo[2], 10, 32)
		if tmp1 == len(resultStr)-1 {
			
		}
	}

	resultStr = strA + strB + strC
	result, _ := strconv.ParseFloat(resultStr, 64)

	return result
}

func calculateStr(inputStr string) string {
	strArr := generateArr(strings.Replace(inputStr, " ", "", -1))

	for i := 0; i < len(strArr); i++ {
		if strArr[i] == "*" || strArr[i] == "/" {
			arrTmp := []string{strArr[i-1], strArr[i], strArr[i+1]}
			tmp := calculateMultipl(arrTmp)
			strArr = removeArr(strArr, i-1)
			strArr = removeArr(strArr, i-1)
			strArr = removeArr(strArr, i-1)
			strArr = addArrIdx(strArr, tmp, i-1)
			i--
		}
	}

	for i := 0; i < len(strArr); i++ {
		if strArr[i] == "+" || strArr[i] == "_" {
			arrTmp := []string{strArr[i-1], strArr[i], strArr[i+1]}
			tmp := calculateAddition(arrTmp)
			strArr = removeArr(strArr, i-1)
			strArr = removeArr(strArr, i-1)
			strArr = removeArr(strArr, i-1)
			strArr = addArrIdx(strArr, tmp, i-1)
			i--
		}
	}

	return strArr[0]
}

func calculateMultipl(strArr []string) string {
	if strArr[1] == "*" {
		tmp1, _ := strconv.ParseFloat(strArr[0], 64)
		tmp2, _ := strconv.ParseFloat(strArr[2], 64)
		result := tmp1 * tmp2

		fmt.Println(result)
		return strconv.FormatFloat(result, 'f', -1, 64)
	} else if strArr[1] == "/" {
		tmp1, _ := strconv.ParseFloat(strArr[0], 64)
		tmp2, _ := strconv.ParseFloat(strArr[2], 64)
		result := tmp1 / tmp2

		fmt.Println(result)
		return strconv.FormatFloat(result, 'f', -1, 64)
	}
	return ""
}

func calculateAddition(strArr []string) string {
	if strArr[1] == "+" {
		tmp1, _ := strconv.ParseFloat(strArr[0], 64)
		tmp2, _ := strconv.ParseFloat(strArr[2], 64)
		result := tmp1 + tmp2

		fmt.Println(result)
		return strconv.FormatFloat(result, 'f', -1, 64)
	} else if strArr[1] == "-" {
		tmp1, _ := strconv.ParseFloat(strArr[0], 64)
		tmp2, _ := strconv.ParseFloat(strArr[2], 64)
		result := tmp1 - tmp2

		fmt.Println(result)
		return strconv.FormatFloat(result, 'f', -1, 64)
	}
	return ""
}

func checkHighestLevel(inputStr string) []string {
	result := []string{"", "", ""}
	//strArr := []rune(inputStr)
	currentLevel := 0
	highestLevel := 0
	startIdx := 0
	endIdx := 0

	for i := 0; i < len(inputStr); i++ {
		if inputStr[i] == '(' {
			currentLevel++
			if currentLevel >= highestLevel {
				highestLevel = currentLevel
				startIdx = i
			}
		}
		if inputStr[i] == ')' {
			if currentLevel == highestLevel {
				endIdx = i
			}
			currentLevel--
		}
	}

	if endIdx == 0 {
		result[0] = inputStr
		result[1] = "0"
		result[2] = strconv.FormatInt(int64(len(inputStr))-1, 32)
	} else {
		result[0] = inputStr[startIdx+1 : endIdx]
		result[1] = strconv.FormatInt(int64(startIdx), 10)
		result[2] = strconv.FormatInt(int64(endIdx), 10)
	}

	return result
}

func generateArr(inputStr string) []string {
	var result []string
	strTmp := string(inputStr[0])

	for i := 1; i < len(inputStr); i++ {
		if inputStr[i] == '-' && (inputStr[i-1] == '*' || inputStr[i-1] == '/' || inputStr[i-1] == '+' || inputStr[i-1] == '-') {
			strTmp = strTmp + string(inputStr[i])
		} else if inputStr[i] == '*' || inputStr[i] == '/' || inputStr[i] == '+' || inputStr[i] == '-' {
			if strTmp != "" {
				result = addArr(result, strTmp)
			}
			result = addArr(result, string(inputStr[i]))
			strTmp = ""
		} else {
			strTmp = strTmp + string(inputStr[i])
		}
	}
	result = addArr(result, strTmp)
	return result
}

func removeArr(inputArr []string, idx int) []string {
	var result []string
	result = append(inputArr[:idx], inputArr[idx+1:]...)
	return result
}

func addArr(inputArr []string, str string) []string {
	var result []string
	result = append(inputArr, str)
	return result
}

func addArrIdx(inputArr []string, str string, idx int) []string {
	var result []string
	var inputArrTmp []string
	var resultTmp []string
	copy(inputArrTmp, inputArr)

	resultTmp = append(inputArrTmp[:idx], str)
	result = append(resultTmp, inputArr[idx:]...)
	return result
}
