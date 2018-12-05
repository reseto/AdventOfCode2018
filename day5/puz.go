package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	fileScanner := bufio.NewScanner(input)
	lns := make(map[int]string) // it's a Set
	i := int(0)
	for fileScanner.Scan() {
		text := fileScanner.Text()
		lns[i] = text
		i++
	}
	lines := make([]string, i)
	for k, v := range lns {
		lines[k] = v
	}
	//sort.Strings(lines)
	//printSorted(lines)
	var proc string
	for z := range lines {
		line := lines[z]
		proc = poly(line[:], 0)
		fmt.Println(proc)
	}
	fmt.Println("part1 ", len(proc))

	min := int(60000)
	for _,c:= range "abcdefghijklmnopqrstuvwxyz" {
		line := lines[0]
		fmt.Println(len(line))
		line = strings.Replace(line, string(c), "", -1)
		line = strings.Replace(line, other(string(c)), "", -1)
		//fmt.Println(len(line))
		proc = poly(line[:], 0)
		if len(proc) < min {
			min = len(proc)
			//fmt.Println("new minimum: ", min)
		}
	}
	fmt.Println("part2 ", min)
}

func poly(inp string, start int) string {
	proc := inp
	var i int
	for i =start; i <len(inp); i++ {
		//fmt.Println(start)
		c := string(inp[i])
		if i>0 {
			a := string(inp[i-1])
			//fmt.Println("converting ", len(proc))
			if a==other(c) {
				proc = proc[:i-1]+proc[i+1:]
				return poly(proc, i-1)
			}
			//fmt.Println("to ", len(proc))
		}
		if i<len(inp)-1 {
			e := string(inp[i+1])
			//fmt.Println("converting ", len(proc))
			if e==other(c) {
				proc = proc[:i]+proc[i+2:]
				return poly(proc, i)
			}
			//fmt.Println("to ", len(proc))
		}
	}
	//fmt.Println(len(proc))
	if len(inp) == len(proc) { return inp }
	return poly(proc, 0)
}

func other(c string) string {
	if strings.ToLower(c) == c { return strings.ToUpper(c) }
	return strings.ToLower(c)
}

func printSorted(lines []string) {
	output, _ := os.OpenFile("sorted.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	for l := range lines {
		_, _ = fmt.Fprintln(output, lines[l])
	}
	_ = output.Close()
}

func mapToArr() {
	m := map[string]string{"1": "a", "2": "b"}
	arr := []string{}
	for k, v := range m {
		arr = append(arr, k, v)
	}
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
