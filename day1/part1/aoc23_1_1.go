package main

import (
    "bufio"
    "fmt"
    "os"
    "unicode"
    "strconv"
)

func main() {
    var first, last rune = rune(0), rune(0)
    var sum int = 0

    file, err := os.Open("aoc.in")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        first = rune(0)
        for _, c := range line {
            if unicode.IsDigit(c) {
                last = c
                if first == rune(0) {
                    first = c
                }
            }
        }
        lineValue, _ := strconv.Atoi(string(first) + string(last))
        sum = sum + lineValue
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    fmt.Println(sum)
}

