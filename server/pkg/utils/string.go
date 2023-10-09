package utils

import (
	"regexp"
)

func RemoveDuplicate[T string | int](sliceList []T) []T {
	uniqueMap := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := uniqueMap[item]; !value {
			uniqueMap[item] = true
			list = append(list, item)
		}
	}
	return list
}

func RemoveTextWithBrackets(s string) string {
	pattern := `\[.+?\]`
	re := regexp.MustCompile(pattern)
	output := re.ReplaceAllString(s, "")
	return output
}
