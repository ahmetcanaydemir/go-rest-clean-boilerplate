package helper

import "strings"

func ContainsFloat64(elements []float64, target float64) bool {
	for _, elem := range elements {
		if elem == target {
			return true
		}
	}
	return false
}

func ContainsString(elements []string, target string) bool {
	for _, elem := range elements {
		if elem == target {
			return true
		}
	}
	return false
}

func SplitByPipe(strList []string) []string {
	if len(strList) != 0 {
		strList = strings.Split(strList[0], "|")
	}
	return strList
}

func SplitByPipeToUpper(strList []string) []string {
	if len(strList) != 0 {
		str := strings.ToUpper(strList[0])
		strList = strings.Split(str, "|")
	}
	return strList
}
