package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Pariatur nisi nostrud id ipsum esse minim velit Lorem eiusmod est reprehenderit sint Esse non proident labore dolore labore nisi dolore dolor proident voluptate minim laborum Ad irure ut Lorem exercitation Ipsum reprehenderit consequat veniam non reprehenderit ut eiusmod magna magna aliquip ut sint Mollit irure fugiat nostrud non excepteur aliquip cillum Labore dolore deserunt irure ea temporConsequat ea quis ipsum esse minim reprehenderit amet Qui do dolor do anim occaecat culpa commodo sunt mollit excepteur laboris tempor Ullamco labore quis culpa laborum sunt Culpa incididunt ad consectetur ut deserunt tempor proident ut Qui quis aliqua fugiat sunt dolore commodo reprehenderit veniam tempor Voluptate sit laboris sunt et do sint ea irure veniam duis laborum Quis irure id officia inNisi velit proident cupidatat laborum velit enim qui deserunt consectetur Ad ea quis veniam cupidatat tempor sunt elit velit Sit anim irure sunt ut exercitation excepteur elit laboris occaecat dolor Duis ad commodo ut quis nisi pariatur anim Fugiat ad id anim velit labore"

	//find the wordsize that occurs the most by size and print the word size and the number of occurrence
	//Hint : use strings.Split()

	//example output : 4 letters occur the most with 29 occurrences
	words := strings.Split(str, " ")
	wordCountBySize := make(map[int]int)
	for _, word := range words {
		wordCountBySize[len(word)]++
	}
	maxSize := 0
	maxCount := 0
	for size, count := range wordCountBySize {
		if count > maxCount {
			maxCount = count
			maxSize = size
		}
	}
	fmt.Printf("%d letter words occur the most with %d occurrances\n", maxSize, maxCount)
}
