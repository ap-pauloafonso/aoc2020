package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func process(str string) int {
	arr := make([]int, 0, 128)
	for a := 0; a <= 127; a++ {
		arr = append(arr, a)
	}
	counter := len(arr)
	for i := 0; i < 7; i++ {
		counter = counter / 2
		if str[i] == byte('F') {
			arr = arr[:counter]
		} else {
			arr = arr[counter:]
		}
	}

	arr2 := make([]int, 0, 8)
	for a := 0; a <= 7; a++ {
		arr2 = append(arr2, a)
	}
	counter2 := len(arr2)
	for j := 7; j < 10; j++ {
		counter2 = counter2 / 2
		if str[j] == byte('L') {
			arr2 = arr2[:counter2]
		} else {
			arr2 = arr2[counter2:]
		}
	}
	return arr[0]*8 + arr2[0]

}
func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(dat), "\n")
	max := -1
	for _, line := range lines {
		r := process(line)
		if r > max {
			max = r
		}
	}
	fmt.Println(max)
}
