package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getList(filename string) (list []string) {
	b, _ := ioutil.ReadFile(filename)

	lines := strings.Split(string(b), "\n")
	list = make([]string, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		list = append(list, l)
	}
	return list
}

func runDiagnostic(oresult []string, pos int) (int, int) {
	ones := 0
	zeros := 0
	for _, c := range oresult {
		if c[pos] == '1' {
			ones++
		} else {
			zeros++
		}
	}
	return ones, zeros
}

func rate(list []string) ([]string, []string) {
	size := len(list[0])
	oxygenList := list
	co2List := list
	for position := 0; position < size; position++ {
		oneCount, zeroCount := runDiagnostic(oxygenList, position)
		if len(oxygenList) > 1 {
			var n []string
			for index, binaryNumber := range oxygenList {
				if ((oneCount >= zeroCount) && binaryNumber[position] == '1') || ((zeroCount > oneCount) && binaryNumber[position] == '0') {
					if index < len(oxygenList) {
						n = append(n, binaryNumber)
					}
				}
			}
			oxygenList = n
		}
		oneCount, zeroCount = runDiagnostic(co2List, position)
		if len(co2List) > 1 {
			var n []string
			for i, c := range co2List {
				if ((oneCount >= zeroCount) && c[position] == '0') || ((zeroCount > oneCount) && c[position] == '1') {
					if i < len(co2List) {
						n = append(n, c)
					}
				}
			}
			co2List = n
		}

	}
	return oxygenList, co2List
}

func main() {
	lines := getList("input")
	oxygen, co2 := rate(lines)
	oxygenInt, _ := strconv.ParseInt(oxygen[0], 2, 64)
	co2Int, _ := strconv.ParseInt(co2[0], 2, 64)
	fmt.Println(oxygenInt * co2Int)
}
