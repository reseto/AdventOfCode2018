package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
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
	sort.Strings(lines)
	//printSorted(lines)

	layout := "2006-01-02 15:04"
	guards := make(map[int]Guard)

	id := int(0)
	for z := range lines {
		line := lines[z]
		words := strings.Fields(line) // split by whitespace
		isGuard := strings.Contains(line, "Guard")
		if isGuard {
			idx, _ := strconv.ParseInt(words[3][1:], 10, 64)
			id = int(idx)
		}
		isFalls := strings.Contains(line, "falls")
		if isFalls {
			start, _ := time.Parse(layout, lines[z][1:1+len(layout)])
			end, _ := time.Parse(layout, lines[z+1][1:1+len(layout)])
			//fmt.Println(id, " asleep from: ", start.Minute(), " to ", end.Minute())

			// MAP:
			if _, exists := guards[id]; !exists {
				//GO IS FUCKING WEIRD, this check for existing value in map is like WTF !
				guards[id] = Guard{id, make([]int, 60), 0, 0}
			}
			for m := start.Minute(); m < end.Minute(); m++ {
				// can't do "guards[id].asleeps[m]++" WTF
				grd := guards[id]
				grd.asleeps[m] ++
			}
		}
	}
	gg := guards[0]
	for g := range guards {
		guard := guards[g]
		max := 0
		for k:=0;k<60;k++ {
			km := guard.asleeps[k]
			if km > max {
				max = km
				guard.mtIndex = k
			}
			guard.totalAsleep += km
		}
		if guard.totalAsleep > gg.totalAsleep {
			gg = guard
		}
		//fmt.Println(guard)
	}
	fmt.Println("part1 ", gg.id,gg.mtIndex,int(gg.id*gg.mtIndex))
	max := 0
	idx := 0
	for g := range guards {
		guard := guards[g]
		for k:=0;k<60;k++ {
			km := guard.asleeps[k]
			if km > max {
				max = km
				gg = guard
				idx = k
			}
		}
	}
	fmt.Println("part2 ", gg.id,idx,int(gg.id*idx))
}

type Guard struct {
	//time       time.Time
	id          int
	asleeps     []int
	totalAsleep int
	mtIndex     int
}

// graveyard

type GList []Guard

func (p GList) Len() int           { return len(p) }
func (p GList) Less(i, j int) bool { return p[i].totalAsleep < p[j].totalAsleep }
func (p GList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

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
