package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
)

const (
	workersParse      = 2
	workersProcessing = 2
	channelBufferSize = 4
)

type inputParsed struct {
	min        int
	max        int
	charactere rune
	text       string
}

func parse() (parsedCh <-chan inputParsed, WaitGroup *sync.WaitGroup) {
	dat, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(dat), "\n")

	var wg sync.WaitGroup
	wg.Add(len(lines))
	inputCH := make(chan string, channelBufferSize)
	outputCH := make(chan inputParsed, channelBufferSize)

	go func() {
		for i := 0; i < workersParse; i++ {
			go workerParse(inputCH, outputCH)
		}
	}()

	go func() {
		for _, line := range lines {
			inputCH <- line
		}
		close(inputCH)
	}()

	return outputCH, &wg
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

	part1()
	part2()
}

func part1() {
	items, wg := parse()
	result := process(items, wg, workerPart1)
	fmt.Println("Part1: ", result)

}
func part2() {
	items, wg := parse()
	result := process(items, wg, workerPart2)
	fmt.Println("Part2: ", result)
}
func process(inputParsedChn <-chan inputParsed, wg *sync.WaitGroup, workerFn func(inChchan <-chan inputParsed, outCh chan<- bool, wg *sync.WaitGroup)) int {
	finalOutputCH := make(chan bool, channelBufferSize)

	go func() {
		for i := 0; i < workersProcessing; i++ {
			go workerFn(inputParsedChn, finalOutputCH, wg)
		}
	}()

	go func() {
		wg.Wait()
		close(finalOutputCH)
	}()

	var counter int
	for boolResult := range finalOutputCH {
		if boolResult {
			counter++
		}
	}
	return counter
}

func workerPart1(in <-chan inputParsed, out chan<- bool, wg *sync.WaitGroup) {
	for v := range in {
		out <- calculatePart1(v)
		wg.Done()
	}
}
func workerPart2(in <-chan inputParsed, out chan<- bool, wg *sync.WaitGroup) {
	for v := range in {
		out <- calculatePart2(v)
		wg.Done()
	}
}

func calculatePart1(input inputParsed) bool {
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

func calculatePart2(input inputParsed) bool {
	first := []rune(input.text)[input.min-1]
	second := []rune(input.text)[input.max-1]

	if first == input.charactere && second != input.charactere || first != input.charactere && second == input.charactere {
		return true
	}
	return false
}
