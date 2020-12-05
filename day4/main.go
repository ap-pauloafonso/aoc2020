package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func validPart2(passwords map[string]string) bool {
	v1, ok1 := passwords["byr"]
	v2, ok2 := passwords["iyr"]
	v3, ok3 := passwords["eyr"]
	v4, ok4 := passwords["hgt"]
	v5, ok5 := passwords["hcl"]
	v6, ok6 := passwords["ecl"]
	v7, ok7 := passwords["pid"]
	// _, ok8 := passwords["cid"]
	return ok1 && validByr(v1) &&
		ok2 && validIyr(v2) &&
		ok3 && validEyr(v3) &&
		ok4 && validHgt(v4) &&
		ok5 && validHcl(v5) &&
		ok6 && validEcl(v6) &&
		ok7 && validPid(v7)
}

func validPart1(passwords map[string]string) bool {
	_, ok1 := passwords["byr"]
	_, ok2 := passwords["iyr"]
	_, ok3 := passwords["eyr"]
	_, ok4 := passwords["hgt"]
	_, ok5 := passwords["hcl"]
	_, ok6 := passwords["ecl"]
	_, ok7 := passwords["pid"]
	// _, ok8 := passwords["cid"]
	return ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && ok7
}
func validByr(str string) bool {
	v, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return v >= 1920 && v <= 2002
}
func validIyr(str string) bool {
	v, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return v >= 2010 && v <= 2020
}
func validEyr(str string) bool {
	v, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}
	return v >= 2020 && v <= 2030
}
func validHgt(str string) bool {
	if strings.HasSuffix(str, "cm") {
		v, err := strconv.Atoi(string([]rune(str)[:len(str)-2]))
		if err != nil {
			panic(err)
		}
		return v >= 150 && v <= 193
	} else if strings.HasSuffix(str, "in") {
		v, err := strconv.Atoi(string([]rune(str)[:len(str)-2]))
		if err != nil {
			panic(err)
		}
		return v >= 59 && v <= 76
	} else {
		return false
	}

}
func validHcl(str string) bool {
	r := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	return r.MatchString(str)
}
func validEcl(str string) bool {
	return str == "amb" ||
		str == "blu" ||
		str == "brn" ||
		str == "gry" ||
		str == "grn" ||
		str == "hzl" ||
		str == "oth"
}

func validPid(str string) bool {
	r := regexp.MustCompile(`^[0-9]{9}$`)
	ret := r.MatchString(str)
	return ret
}

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	linesReg := regexp.MustCompile("\n\n")
	lines := linesReg.Split(string(dat), -1)
	itemsReg := regexp.MustCompile(` |\n`)
	var counterPart1 int
	var counterPart2 int

	for _, line := range lines {
		result := itemsReg.Split(line, -1)
		m := make(map[string]string)
		for _, v := range result {
			splittedResult := strings.Split(v, ":")
			m[splittedResult[0]] = splittedResult[1]
		}
		if validPart1(m) {
			counterPart1++
		}
		if validPart2(m) {
			counterPart2++
		}
	}

	fmt.Println("Part1: ", counterPart1)
	fmt.Println("Part2: ", counterPart2)
}
