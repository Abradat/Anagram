package main

import (
	"fmt"
	"bufio"
	"os"

)

func main() {

	var dictionary = make(map[string][]string)

	dictionary = generateAnagrams()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("You can Do 2 things\n\n1)Anagram\n2)EXIT\n\nEnter your Decision: ")
		text, _ := reader.ReadString('\n')
		if(text[:len(text) - 1] == "1") {
			fmt.Print("Enter The word for anagram: ")
			text, _ := reader.ReadString('\n')
			results := dictionary[SortString(text[:len(text) - 1])]
			if(len(results) == 0) {
				fmt.Printf("The is NO Result !!!\n")

			} else {
				fmt.Printf("The results are these : \n")
				for cnt, result := range results {
					fmt.Printf("%d) %s\n", cnt + 1, result)
				}
				fmt.Printf("\n***********\n\n")
			}
		} else {
			fmt.Printf("BYE BYE :)\n")
			break
		}
	}


}
