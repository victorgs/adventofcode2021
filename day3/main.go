package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func getList(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("%s", err)
	}
	lines := strings.Split(string(content), "\n")
	return lines
}

func main() {
	lines := getList("input")
	zeroCount := 0
	oneCount := 0
	var gamma string
	var epsilon string
	var num string
	binLength := len(lines[0])
	pos1 := 0
	pos2 := 1
	for i := 0; i <= binLength-1; i++ {
		zeroCount = 0
		oneCount = 0
		for _, line := range lines {
			num = line[pos1:pos2]
			if num == "0" {
				zeroCount = zeroCount + 1
			}
			if num == "1" {
				oneCount = oneCount + 1
			}

		}
		if zeroCount < oneCount {
			gamma += "1"
			epsilon += "0"
		}
		if zeroCount > oneCount {
			gamma += "0"
			epsilon += "1"
		}
		pos1 = pos1 + 1
		pos2 = pos2 + 1
	}
	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Println(gammaInt * epsilonInt)
}
