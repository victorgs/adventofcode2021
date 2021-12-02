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
	depth := 0
	depthIncrease := 0
	horizontalPos := 0
	aim := 0
	lines := getList("input")
	var value int
	for _, line := range lines {
		split := strings.Split(line, " ")
		value, _ = strconv.Atoi(split[1])
		switch split[0] {
		case "forward":
			horizontalPos = horizontalPos + value
			depthIncrease = value * aim
			depth = depth + depthIncrease
		case "down":
			aim = aim + value
		case "up":
			aim = aim - value
		}
	}
	total := depth * horizontalPos
	fmt.Println(total)
}
