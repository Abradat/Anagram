package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"os"
)

func unique(StringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range StringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}


func generateAnagrams() map[string][]string {

	fmt.Printf("\n\n\n******* Going to Start generating anagrams *******\n\n\n")
	var finalDic = make(map[string][]string)
	var dic = make(map[string][]string)
	gopath := os.Getenv("GOPATH")
	//fmt.Printf("%s\n\n", gopath)
	excelFileName := gopath + "/src/anagram/models/PersianWords.xlsx"
	xlFile, _ := xlsx.OpenFile(excelFileName)

	counter := 0

	sheet := xlFile.Sheets[0]
	for _, row := range sheet.Rows {
		for _, cell := range row.Cells {
			text := cell.String()
			sortedText := SortString(text)
			counter += 1
			fmt.Printf("%d\n", counter)
			//fmt.Printf("id : %d is %s\n", cnt, SortString(text))
			dic[sortedText] = append(dic[sortedText], text)


		}
	}


	for key, value := range dic {

		finalValue := unique(value)
		finalDic[key] = finalValue
		//	fmt.Printf("key is : %s and the value is %v\n\n", key, finalValue)
	}
	fmt.Printf("******* Done *******\n\n")

	return finalDic

}
