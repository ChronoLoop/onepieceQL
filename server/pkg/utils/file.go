package utils

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
)

func SaveCSVFile(in interface{}, csvPath string) {
	tempFilePath := "./data/temp.csv"
	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return
	}
	defer tempFile.Close()
	if err := gocsv.MarshalFile(in, tempFile); err != nil {
		panic(err)
	}
	err = os.Rename(tempFilePath, csvPath)
	if err != nil {
		fmt.Println("Error replacing the original file:", err)
		return
	}
}
