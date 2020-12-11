package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type instuction struct {
	op  string
	arg int
}

func parse() []instuction {

	f, _ := os.Open("input.txt")

	bytes, _ := ioutil.ReadAll(f)

	items := make([]instuction, 0)
	for _, v := range strings.Split(string(bytes), "\n") {
		re := regexp.MustCompile(`(nop|acc|jmp)\s([+-]\d+)`)
		ret := re.FindStringSubmatch(v)
		n, _ := strconv.Atoi(ret[2])
		items = append(items, instuction{op: ret[1], arg: n})
	}
	return items
}
func result1(instructions []instuction) (acumulator int, halt error) {
	currentLine := 0
	accumulator := 0
	visited := make(map[int]struct{}, 0)
	for {
		if currentLine > len(instructions)-1 {
			break
		}
		ins := instructions[currentLine]
		if _, ok := visited[currentLine]; ok {
			return accumulator, errors.New("halt")
		}
		visited[currentLine] = struct{}{}

		switch ins.op {
		case "nop":
			{
				currentLine++
			}
		case "acc":
			{
				currentLine++
				accumulator += ins.arg
			}
		case "jmp":
			{
				currentLine = currentLine + ins.arg
			}
		}
	}
	return accumulator, nil
}

func result2(instructions []instuction) int {
	for i := len(instructions) - 1; i >= 0; i-- {
		if instructions[i].op == "nop" || instructions[i].op == "jmp" {
			originalCopy := instructions[i]
			if instructions[i].op == "nop" {
				instructions[i].op = "jmp"
			} else {
				instructions[i].op = "nop"
			}
			result, haltError := result1(instructions)
			if haltError == nil {
				return result
			}
			instructions[i] = originalCopy
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
