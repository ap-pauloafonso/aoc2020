package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type lineObj struct {
	columnCount   int
	lineNumber    int
	iteration     int
	characteres   string
	hasTree       bool
	resultPartStr string
}

func newLine(line, currentColumn int, characteres string) lineObj {
	return lineObj{columnCount: currentColumn, lineNumber: line, characteres: characteres}
}

func process(lines []string, right, down int) int {
	arr := make([]lineObj, 0, len(lines))
	for idx, line := range lines {
		arr = append(arr, newLine(idx, 0, line))
	}
	textLength := len(arr[0].characteres)
	currentColumn := 0
	// arr[0].resultPartStr = arr[0].characteres
	for i := 0; i < len(lines)-down; i = i + down {
		currentColumn += right
		arr[i+down].columnCount = currentColumn
		arr[i+down].iteration = int(currentColumn / textLength)

		if rune(arr[i+down].characteres[currentColumn-arr[i+down].iteration*textLength]) == rune('#') {
			// arr[i+down].resultPartStr = string(arr[i+down].characteres[:currentColumn-arr[i+down].iteration*textLength]) + "X" + string(arr[i+down].characteres[currentColumn-arr[i+down].iteration*textLength+down:])
			arr[i+down].hasTree = true
			continue
		}
		// arr[i+down].resultPartStr = string(arr[i+down].characteres[:currentColumn-arr[i+down].iteration*textLength]) + "O" + string(arr[i+down].characteres[currentColumn-arr[i+down].iteration*textLength+down:])
		arr[i+down].hasTree = false
	}

	var counter int
	for _, v := range arr {
		// fmt.Printf("%+v\n", v)
		if v.hasTree {
			counter++
		}
	}

	return counter
}
func part1(lines []string) {
	fmt.Println("Part1: ", process(lines, 3, 1))
}

func part2(lines []string) {
	a := process(lines, 1, 1)
	b := process(lines, 3, 1)
	c := process(lines, 5, 1)
	d := process(lines, 7, 1)
	e := process(lines, 1, 2)

	// fmt.Println(a, b, c, d, e)

	fmt.Println("Part2: ", a*b*c*d*e)
}

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(dat), "\r\n")
	part1(lines)
	part2(lines)
}
