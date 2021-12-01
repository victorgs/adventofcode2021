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
	var increases int
	var buffer []int
	for _, line := range lines {
		current, _ := strconv.Atoi(line)
		buffer = append(buffer, current)
		if len(buffer) >= 4 {
			windowA := buffer[0] + buffer[1] + buffer[2]
			windowB := buffer[1] + buffer[2] + buffer[3]
			if windowB > windowA {
				increases++
			}
			buffer = buffer[1:]
		}
	}
	fmt.Println(increases)
}
