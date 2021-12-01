package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
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
	var increases int = 0
	var prev int = math.MaxInt64
	for _, line := range lines {
		current, _ := strconv.Atoi(line)
		if current > prev {
			increases++
		}
		prev = current
	}
	fmt.Println(increases)
}
