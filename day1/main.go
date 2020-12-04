package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	dat, _ := ioutil.ReadFile("input.txt")
	input := string(dat)
	x, y := part1(input)
	x2, y2, z := part2(input)
	fmt.Println("Part 1 result: ", x*y)
	fmt.Println("Part 2 result: ", x2*y2*z)

}

func part1(input string) (xResult int, YResult int) {
	numbers := strings.Split(input, "\n")
	for i := 0; i <= len(numbers)-1; i++ {
		for j := 0; j <= len(numbers)-1; j++ {
			jint, _ := strconv.Atoi(numbers[j])
			iint, _ := strconv.Atoi(numbers[i])
			if (jint + iint) == 2020 {
				xResult, YResult = jint, iint
				return
			}
		}
	}
	panic("fail")
}

func part2(input string) (xResult, YResult, zResult int) {
	numbers := strings.Split(input, "\n")
	for i := 0; i <= len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-1; j++ {
			for k := 0; k <= len(numbers)-1; k++ {
				jint, _ := strconv.Atoi(numbers[j])
				iint, _ := strconv.Atoi(numbers[i])
				kint, _ := strconv.Atoi(numbers[k])
				if (jint + iint + kint) == 2020 {
					xResult, YResult, zResult = jint, iint, kint
					return
				}
			}
		}
	}
	panic("fail")
}
