package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

var found = false
var sum = int(0)
var set = make(map[int]bool)

func main() {

    max := 0
    min := 0

    fileHandle, _ := os.Open("puzzle1.input")
    defer fileHandle.Close()

    for ; !found; _, _ = fileHandle.Seek(0, 0) {
        fileScanner := bufio.NewScanner(fileHandle)
        for fileScanner.Scan() {
            text := fileScanner.Text()
            j, _ := strconv.ParseInt(text, 10, 64)
            i := int(j)
            sum += i
            if sum > max {
                max = sum
            }
            if sum < min {
                min = sum
            }
            search()
        }
    }
    fmt.Println("sum: ", sum)
    // for fun stats
    fmt.Println("max: ", max)
    fmt.Println("min: ", min)
}

func search() {
    if !found {
        if !set[sum] {
            set[sum] = true
        } else {
            found = true
            fmt.Println("first duplicate: ", sum)
        }
    }
}

// graveyard

//if text[0]=='-' {
//    sum -= text[1:]
//} else if text[0]=='+' {
//
//} else {
//    fmt.Println("error parsing ", text)
//}