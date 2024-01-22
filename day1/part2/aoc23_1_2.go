package main

import (
    "bufio"
    "fmt"
    "os"
    "unicode"
    "strconv"
    "strings"
)

func main() {
    var numbers = map[string]int {
        "one":   1,
        "two":   2,
        "three": 3,
        "four":  4,
        "five":  5,
        "six":   6,
        "seven": 7,
        "eight": 8,
        "nine":  9,
    }

    var first, last int = -1, -1
    var sum int = 0

    file, err := os.Open("aoc.in")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        first = -1
        for i, c := range line {
            for k, v := range numbers {
                if _, found := strings.CutPrefix(line[i:], k); found {
                    last = v
                    if first == -1 {
                        first = v
                    }
                    break
                }
            }
            if unicode.IsDigit(c) {
                last, _ = strconv.Atoi(string(c))
                if first == -1 {
                    first, _ = strconv.Atoi(string(c))
                }
            }
        }
        // fmt.Println(line)
        // fmt.Println(first)
        // fmt.Println(last)
        lineValue := first * 10 + last
        sum = sum + lineValue
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    fmt.Println(sum)
}

