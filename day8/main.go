package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type instuction struct {
	op  string
	arg int
}

func parse() []*instuction {

	f, _ := os.Open("input.txt")

	bytes, _ := ioutil.ReadAll(f)

	items := make([]*instuction, 0)
	for _, v := range strings.Split(string(bytes), "\n") {
		splits := strings.Split(v, " ")
		op := splits[0]
		n, _ := strconv.Atoi(splits[1])
		items = append(items, &instuction{op: op, arg: n})
	}
	return items
}
func result1(instructions []*instuction) (acumulator int, halt bool) {
	currentLine := 0
	accumulator := 0
	visited := make(map[int]struct{}, 0)
	for {
		if currentLine > len(instructions)-1 {
			break
		}
		ins := instructions[currentLine]
		if _, ok := visited[currentLine]; ok {
			return accumulator, true
		}
		visited[currentLine] = struct{}{}

		if ins.op == "nop" {
			currentLine++
			continue
		}
		if ins.op == "acc" {
			currentLine++
			accumulator += ins.arg
			continue
		}
		if ins.op == "jmp" {
			currentLine = currentLine + ins.arg
			continue
		}

	}
	return accumulator, false
}

func result2(instructions []*instuction) int {
	for i := len(instructions) - 1; i >= 0; i-- {
		if instructions[i].op == "nop" || instructions[i].op == "jmp" {
			originalCopy := *instructions[i]
			if instructions[i].op == "nop" {
				instructions[i].op = "jmp"
			} else {
				instructions[i].op = "nop"
			}
			result, halt := result1(instructions)
			if !halt {
				return result
			}
			instructions[i] = &originalCopy
		}
	}
	panic("all combinations halt")
}

func main() {
	instructions := parse()
	r1, _ := result1(instructions)
	fmt.Println("Result1: ", r1)
	fmt.Println("Result2: ", result2(instructions))

}
