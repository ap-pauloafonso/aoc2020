package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	preamble = 25
)

func parse() []int {
	f, _ := os.Open("input.txt")
	bytes, _ := ioutil.ReadAll(f)
	ret := make([]int, 0)
	s := strings.Split(string(bytes), "\r\n")
	for _, v := range s {
		n, _ := strconv.Atoi(v)
		ret = append(ret, n)
	}
	return ret
}
func sumMap(numbers []int) map[int]struct{} {
	ret := make(map[int]struct{}, 0)
	for i := 0; i < preamble; i++ {
		for j := 0; j < preamble; j++ {
			if i != j {
				ret[numbers[i]+numbers[j]] = struct{}{}
			}
		}
	}
	return ret
}

func result1(numbers []int) int {
	current := preamble
	for {
		preambleSumMap := sumMap(numbers[current-preamble : current])
		if _, ok := preambleSumMap[numbers[current]]; !ok {
			return numbers[current]
		}
		current++
	}
}
func result2(target int, numbers []int) int {
	lowerIdx := 0
	higherIdx := 0

	offset := 0
	for {
		current := offset
		if current > len(numbers)-1 {
			break
		}
		counter := numbers[current]
		for {
			if current+1 > len(numbers)-1 || counter+numbers[current+1] > target {
				break
			} else if counter+numbers[current+1] == target {
				lowerIdx = offset
				higherIdx = current + 1
				break
			}
			counter += numbers[current+1]
			current++
		}
		if higherIdx > 0 {
			break
		}
		offset++
	}
	filterArr := numbers[lowerIdx : higherIdx+1]
	sort.Ints(filterArr)

	r := filterArr[0] + filterArr[len(filterArr)-1]

	return r
}
func main() {
	numbers := parse()
	fmt.Println("Result1: ", result1(numbers))
	fmt.Println("Result2: ", result2(result1(numbers), numbers))
}
