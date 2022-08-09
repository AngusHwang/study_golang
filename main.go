package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(evaluate(`(1 + 2) + (5 * 3 / (72 - 23))`))
}

func evaluate(expression string) float64 {
	var (
		levelInfo []string
		resultStr string
		strA      string
		strB      string
		strC      string
	)
	i := 0
	resultStr = expression + " "
	for strings.Contains(resultStr, " ") || strings.Contains(resultStr, "(") || strings.Contains(resultStr, ")") || strings.Contains(resultStr, "sin") || strings.Contains(resultStr, "cos") || strings.Contains(resultStr, "tan") {
		levelInfo = checkHighestLevel(resultStr)
		tmp1, _ := strconv.Atoi(levelInfo[1])
		strA = subString(resultStr, 0, tmp1)
		strB = calculateStr(levelInfo[0])
		if i, _ := strconv.Atoi(levelInfo[2]); i == len(resultStr)-1 {
			strC = ""
		} else {
			tmp2, _ := strconv.Atoi(levelInfo[2])
			strC = subString(resultStr, tmp2+1, len(resultStr))
		}
		resultStr = strA + strB + strC
		fmt.Println(strings.Contains(resultStr, " "))
		fmt.Println(strings.Contains(resultStr, "("))
		fmt.Println(strings.Contains(resultStr, ")"))
		fmt.Println(strings.Contains(resultStr, "+"))
		fmt.Println(strings.Contains(resultStr, "-"))
		fmt.Println(strings.Contains(resultStr, "*"))
		fmt.Println(strings.Contains(resultStr, "/"))
		fmt.Println(strings.Contains(resultStr, "sin"))
		fmt.Println(strings.Contains(resultStr, "cos"))
		fmt.Println(strings.Contains(resultStr, "tan"))
		fmt.Println(resultStr)
		fmt.Println(`************************************************** 플로트 스트링변환쪽에서 + 받고오기때문에 무한루프돌음`)

		i++
		if i > 20 {
			break
		}
	}

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
		if strArr[i] == "+" || strArr[i] == "-" {
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

		return strconv.FormatFloat(result, 'f', -1, 64)
	} else if strArr[1] == "/" {
		tmp1, _ := strconv.ParseFloat(strArr[0], 64)
		tmp2, _ := strconv.ParseFloat(strArr[2], 64)
		result := tmp1 / tmp2

		return strconv.FormatFloat(result, 'f', -1, 64)
	}
	return ""
}

func calculateAddition(strArr []string) string {
	if strArr[1] == "+" {
		tmp1, _ := strconv.ParseFloat(strArr[0], 32)
		tmp2, _ := strconv.ParseFloat(strArr[2], 32)
		result := tmp1 + tmp2

		return strconv.FormatFloat(result, 'f', -1, 32)
	} else if strArr[1] == "-" {
		tmp1, _ := strconv.ParseFloat(strArr[0], 32)
		tmp2, _ := strconv.ParseFloat(strArr[2], 32)
		result := tmp1 - tmp2

		return strconv.FormatFloat(result, 'f', -1, 32)
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

/*Utility */
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
	var resultTmp []string
	inputArrTmp := make([]string, len(inputArr))
	copy(inputArrTmp, inputArr)
	resultTmp = append(inputArrTmp[:idx], str)
	result = append(resultTmp, inputArr[idx:]...)
	return result
}

func subString(inputStr string, beginIdx int, endIdx int) string {
	result := inputStr[beginIdx:endIdx]
	return result
}
