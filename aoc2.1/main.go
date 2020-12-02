package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	workers = 4
)

type inputParsed struct {
	min        int
	max        int
	charactere rune
	text       string
}

func parse() []inputParsed {
	dat, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(dat), "\n")

	inputCH := make(chan string, len(lines))
	outputCH := make(chan inputParsed, len(lines))

	go func() {
		for i := 0; i < workers; i++ {
			workerParse(inputCH, outputCH)
		}
	}()
	for _, line := range lines {
		inputCH <- line
	}

	result := make([]inputParsed, 0, len(lines))

	for range lines {
		result = append(result, <-outputCH)
	}

	return result
}

func workerParse(in <-chan string, out chan<- inputParsed) {
	for v := range in {
		out <- procesParse(v)
	}
}

func procesParse(line string) inputParsed {
	min, _ := strconv.Atoi(strings.Split(strings.Split(line, ":")[0], "-")[0])
	max, _ := strconv.Atoi(strings.Split(strings.Split(strings.Split(line, ":")[0], "-")[1], " ")[0])
	charactere := []rune(strings.Split(strings.Split(line, ":")[0], " ")[1])[0]
	text := strings.Split(line, ": ")[1]
	return inputParsed{min: min, max: max, charactere: charactere, text: text}
}

func main() {
	items := parse()
	result := process(items)
	fmt.Println(result)

}
func process(items []inputParsed) int {
	inputCH := make(chan inputParsed, len(items))
	outputCH := make(chan bool, len(items))

	go func() {
		for i := 0; i < workers; i++ {
			worker(inputCH, outputCH)
		}
	}()

	for _, v := range items {
		inputCH <- v
	}
	close(inputCH)

	var counter int
	for range items {
		boolResult := <-outputCH
		if boolResult {
			counter++
		}
	}
	return counter
}

func worker(in <-chan inputParsed, out chan<- bool) {
	for v := range in {
		out <- calculate(v)
	}
}

func calculate(input inputParsed) bool {
	var count int
	for _, c := range input.text {
		if input.charactere == c {
			count++
		}
		if count > input.max {
			return false
		}
	}

	if count >= input.min && count <= input.max {
		return true
	}
	return false
}
