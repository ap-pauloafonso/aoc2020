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
	count int
	child string
}
type node struct {
	name   string
	childs []childNode
}

// type childNode2 struct {
// 	count int
// 	child *node2
// }
// type node2 struct {
// 	name   string
// 	childs []childNode2
// }

func parseLines(lines []string) map[string]*node {
	parentRegex := regexp.MustCompile(`^([a-z]+\s[a-z]+)`)
	childRegex := regexp.MustCompile(`((\d+)\s([a-z]+\s[a-z]+))`)
	nodes := make(map[string]*node)
	for _, line := range lines {
		parentNode := parentRegex.FindStringSubmatch(line)
		var childs []childNode
		for _, group := range childRegex.FindAllStringSubmatch(line, -1) {
			bagName := group[len(group)-1:][0]
			count, _ := strconv.Atoi(string(group[len(group)-2:][0]))
			if _, ok := nodes[bagName]; !ok {
				nodes[bagName] = &node{name: bagName}
			}
			childs = append(childs, childNode{count: count, child: bagName})
		}
		nodes[parentNode[0]] = &node{
			name:   parentNode[0],
			childs: childs,
		}
	}
	return nodes
}

// func parseLines2(lines []string) map[string]*node2 {
// 	parentRegex := regexp.MustCompile(`^([a-z]+\s[a-z]+)`)
// 	childRegex := regexp.MustCompile(`((\d+)\s([a-z]+\s[a-z]+))`)
// 	nodes := make(map[string]*node2)
// 	for _, line := range lines {
// 		parentNode := parentRegex.FindStringSubmatch(line)
// 		var childs []childNode2
// 		for _, group := range childRegex.FindAllStringSubmatch(line, -1) {
// 			bagName := group[len(group)-1:][0]
// 			count, _ := strconv.Atoi(string(group[len(group)-2:][0]))
// 			if _, ok := nodes[bagName]; !ok {
// 				nodes[bagName] = &node2{name: bagName}
// 			}
// 			childs = append(childs, childNode2{count: count, child: nodes[bagName]})
// 		}
// 		nodes[parentNode[0]] = &node2{
// 			name:   parentNode[0],
// 			childs: childs,
// 		}
// 	}
// 	return nodes
// }

func findResults(wantBag string, currentNode *node, visited map[string]bool, nodes map[string]*node) bool {
	var f func(currentNode *node) bool
	f = func(currentNode *node) bool {
		if hasItem, ok := visited[currentNode.name]; ok && !hasItem {
			return false
		}
		if len(currentNode.childs) == 0 {
			return false
		}
		for _, v := range currentNode.childs {
			if v.child == wantBag && v.count > 0 || f(nodes[v.child]) {
				visited[currentNode.name] = true
				return true
			}

		}
		return false
	}
	return f(currentNode)
}

// func findResults2(wantBag string, nodes map[string]*node2) int {
// 	var f func(currentNode *node2) int

// 	f = func(n *node2) int {
// 		var counter int
// 		currentNode := nodes[n.name]
// 		if currentNode.childs == nil || len(currentNode.childs) == 0 {
// 			return 1
// 		}
// 		for _, v := range currentNode.childs {
// 			ret := v.count * f(v.child)
// 			counter += ret
// 		}

// 		return counter
// 	}
// 	result := f(nodes[wantBag])

// 	return result
// }

func main() {
	f, _ := os.Open("input.txt")
	buf, _ := ioutil.ReadAll(f)

	nodes := parseLines(strings.Split(string(buf), "\n"))
	visitedNodes := make(map[string]bool)
	var counter int
	for _, v := range nodes {
		if findResults("shiny gold", v, visitedNodes, nodes) {
			counter++
		}
	}
	fmt.Println("Results1: ", counter)

	////////////////////////////////////////////////////////////////////////
	// nodes2 := parseLines2(strings.Split(string(buf), "\n"))
	// var counter2 int
	// counter2 = findResults2("shiny gold", nodes2)
	// fmt.Println("Results2: ", counter2)
}
