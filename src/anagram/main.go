package main

import "fmt"

func main() {

	var dictionary = make(map[string][]string)

	dictionary = generateAnagrams()

	fmt.Print(dictionary[SortString("قلابت")])
}
