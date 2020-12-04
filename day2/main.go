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

func procesParse(line string) inputParsed {
	splits := strings.Split(line, " ")
	minmax := strings.Split(splits[0], "-")
	min, _ := strconv.Atoi(minmax[0])
	max, _ := strconv.Atoi(minmax[1])
	charactere := []rune(splits[1])[0]
	text := splits[2]
	return inputParsed{min: min, max: max, charactere: charactere, text: text}
}

func workerParse(in <-chan string, out chan<- inputParsed) {
	for v := range in {
		out <- procesParse(v)
	}
}
func parse() (parsedCh <-chan inputParsed, WaitGroup *sync.WaitGroup) {
	dat, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(dat), "\n")

	var wg sync.WaitGroup
	wg.Add(len(lines))
	inputCH := make(chan string, channelBufferSize)
	outputCH := make(chan inputParsed, channelBufferSize)

	for i := 0; i < workersParse; i++ {
		go workerParse(inputCH, outputCH)
	}

	go func() {
		for _, line := range lines {
			inputCH <- line
		}
		close(inputCH)
	}()

	return outputCH, &wg
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
	if input.min-1 < 0 || input.max-1 > len(input.text) {
		return false
	}
	first := []rune(input.text)[input.min-1]
	second := []rune(input.text)[input.max-1]

	if first == input.charactere && second != input.charactere || first != input.charactere && second == input.charactere {
		return true
	}
	return false
}

func process(inputParsedChn <-chan inputParsed, wg *sync.WaitGroup, workerFn func(inChchan <-chan inputParsed, outCh chan<- bool, wg *sync.WaitGroup)) int {
	finalOutputCH := make(chan bool, channelBufferSize)

	for i := 0; i < workersProcessing; i++ {
		go workerFn(inputParsedChn, finalOutputCH, wg)
	}

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

func main() {
	part1()
	part2()
}
