package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"time"
)

func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
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

func PrintStruct(s interface{}) {
	bytes, _ := json.MarshalIndent(s, "", "  ")
	fmt.Printf("\nstruct: %s\n", string(bytes))
}

func GetExePath() string {
	ex, err := os.Executable()

	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	//air runs executable in tmp
	if path.Base(exPath) == "tmp" {
		exPath = filepath.Join(exPath, "../")
	}

	return exPath
}

func RemoveTextWithBrackets(s string) string {
	pattern := `\[.+?\]`
	re := regexp.MustCompile(pattern)
	output := re.ReplaceAllString(s, "")
	return output
}
