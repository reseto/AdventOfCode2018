package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.Open("input")
	defer input.Close()

	fileScanner := bufio.NewScanner(input)
	lines := make(map[int]string)
	i := int(0)
	for fileScanner.Scan() {
		text := fileScanner.Text()
		lines[i] = text
		i++
	}
	var cnt [1000][1000]int
	for i=0;i< len(lines);i++ {
		words := strings.Fields(lines[i])
		x, y := toInt(words[2])
		a, b := toInt(words[3])
		for r := y; r < y+b; r++ {
			for c := x; c < x+a ; c++ {
				cnt[r][c]++
			}
		}
	}
	// part two check for overlap with nothing
	for i=0;i< len(lines);i++ {
		words := strings.Fields(lines[i])
		x, y := toInt(words[2])
		a, b := toInt(words[3])
		isThisIt := true
		out:
		for r := y; r < y+b; r++ {
			for c := x; c < x+a ; c++ {
				if cnt[r][c] != 1 {
					isThisIt = false
					break out
				}
			}
		}
		if isThisIt {
			fmt.Println(lines[i])
		}
	}
	sum := int(0)
	for r := 0; r < 1000; r++ {
		for c := 0; c < 1000 ; c++ {
			if cnt[r][c] > 1 {
				sum ++
			}
		}
		//fmt.Println(cnt[r])
	}
	fmt.Println(sum)


}

func toInt(word string) (int, int) {
	if strings.Contains(word, "x") {
		ss := strings.Split(word, "x")
		a, _ := strconv.ParseInt(ss[0], 10, 64)
		b, _ := strconv.ParseInt(ss[1], 10, 64)
		return int(a),int(b)
	}
	ss := strings.Split(word, ",")
	a, _ := strconv.ParseInt(ss[0], 10, 64)
	y := ss[1]
	b, _ := strconv.ParseInt(y[:len(y)-1], 10, 64)
	return int(a),int(b)
}
