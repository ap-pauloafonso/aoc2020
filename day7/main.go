package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type childNode struct {
	count     int
	childName string
}
type node struct {
	name   string
	childs []childNode
}

func parseLines(lines []string) map[string]node {
	parentRegex := regexp.MustCompile(`^([a-z]+\s[a-z]+)`)
	childRegex := regexp.MustCompile(`((\d+)\s([a-z]+\s[a-z]+))`)
	nodes := make(map[string]node)
	for _, line := range lines {
		parentNode := parentRegex.FindStringSubmatch(line)
		var childs []childNode
		for _, group := range childRegex.FindAllStringSubmatch(line, -1) {
			bagName := group[len(group)-1:][0]
			count, _ := strconv.Atoi(string(group[len(group)-2:][0]))
			if _, ok := nodes[bagName]; !ok {
				nodes[bagName] = node{name: bagName}
			}
			childs = append(childs, childNode{count: count, childName: bagName})
		}
		nodes[parentNode[0]] = node{
			name:   parentNode[0],
			childs: childs,
		}
	}
	return nodes
}

func findResults(wantBag string, nodes map[string]node) int {
	var f func(currentNode node) bool
	visitedNodes := make(map[string]bool)
	f = func(currentNode node) bool {
		if hasItem, ok := visitedNodes[currentNode.name]; ok && !hasItem {
			return false
		}
		if len(currentNode.childs) == 0 {
			return false
		}
		for _, v := range currentNode.childs {
			if v.childName == wantBag && v.count > 0 || f(nodes[v.childName]) {
				visitedNodes[currentNode.name] = true
				return true
			}
		}
		return false
	}

	var counter int
	for _, v := range nodes {
		if f(v) {
			counter++
		}
	}
	return counter
}

func findResults2(wantBag string, nodes map[string]node) int {
	var f func(currentNode node) int
	f = func(n node) int {
		var counter int
		currentNode := nodes[n.name]
		for _, v := range currentNode.childs {
			ret := v.count * (f(nodes[v.childName]) + 1)
			counter += ret
		}
		return counter
	}
	result := f(nodes[wantBag])
	return result
}

func main() {
	f, _ := os.Open("./input.txt")
	buf, _ := ioutil.ReadAll(f)

	nodes := parseLines(strings.Split(string(buf), "\n"))

	fmt.Println("Results1: ", findResults("shiny gold", nodes))
	fmt.Println("Results2: ", findResults2("shiny gold", nodes))
}
