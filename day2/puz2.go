package main

import (
	"bufio"
	"fmt"
	"os"
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

	// iterate length of row words, we will be omitting the r-th character
	for r := int(0); r < len(lines[0]); r++ {
		// iterate all lines
		for a := int(0); a < len(lines)-1; a++ {
			lineA := lines[a]
			modA :=  lineA[:r] + lineA[r+1:]
			// compare with lines not previously visited
			for v := a+1; v < len(lines); v++ {
				lineV := lines[v]
				modV :=  lineV[:r] + lineV[r+1:]
				if modA==modV {
					fmt.Println(modA)
					fmt.Println("omitted character at position:", r)
					fmt.Println("omitted character:", string(lineV[r]))
					fmt.Println("# iterations: ", r*a*v)
					os.Exit(0)
				}
			}
		}
	}

}
// graveyard
//
//func omit(k int, lines map[int]string) map[int]string {
//	mod := make(map[int]string)
//	for i:=int(0);i<len(lines);i++ {
//		//changed my mind to make 3x nested for
//	}
//	return mod
//}