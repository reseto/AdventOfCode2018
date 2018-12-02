package main


import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	two := 0
	tri := 0

	input, _ := os.Open("input")
	defer input.Close()
	//output, _ := os.Open("out")
	//defer output.Close()
	//fmt.Fprint(output, "asdf")

	//for ; !found; _, _ = input.Seek(0, 0) {
		fileScanner := bufio.NewScanner(input)
		for fileScanner.Scan() {
			text := fileScanner.Text()
			cnt := make(map[int]int)
			for i:=int(0); i<26; i++ {
				s := text[i]
				c := int(s)
				fmt.Printf("%4d", c)
				cnt[c] += 1
			}
			fmt.Println()
			base:=97
			for i:=int(0); i<26; i++ {
				if cnt[base+i] == 2 {
					two += 1
					break
				}
			}
			for i:=int(0); i<26; i++ {
				if cnt[base+i] == 3 {
					tri += 1
					break
				}
			}
		}
	//}
	fmt.Println("2: ", two)
	fmt.Println("3: ", tri)
	fmt.Println("checksum: ", two * tri)
}