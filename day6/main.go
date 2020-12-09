package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func answerCountPart1(groupstring string) int {
	foundAnswers := make(map[byte]struct{})
	for i := 0; i < len(groupstring); i++ {
		strByte := groupstring[i]
		if _, ok := foundAnswers[strByte]; !ok && strByte != byte('\n') {
			foundAnswers[strByte] = struct{}{}
		}
	}
	return len(foundAnswers)
}

func answerCountPart2(groupstring string) int {
	peoples := strings.Split(groupstring, "\n")
	var resultCounter int

	foundAnswers := make(map[byte]int)
	for _, people := range peoples {
		for j := 0; j < len(people); j++ {
			strByte := people[j]
			foundAnswers[strByte]++
		}
	}
	for _, v := range foundAnswers {
		if v == len(peoples) {
			resultCounter++
		}
	}

	return resultCounter
}

func main() {
	f, _ := os.Open("input.txt")
	bytes, _ := ioutil.ReadAll(f)
	groups := strings.Split(string(bytes), "\n\n")
	var counter int
	var counterPart2 int
	for _, s := range groups {
		counter += answerCountPart1(s)
		counterPart2 += answerCountPart2(s)
	}
	fmt.Println("result par1: ", counter)
	fmt.Println("result part2: ", counterPart2)

}
