package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func parseCube(cube string) (int, string) {
    items := strings.Split(cube, " ")
    number, _ := strconv.Atoi(items[0])
    color := items[1]
    return number, color
}

func main() {
    var maxCubes = map[string]int {
        "red":   12,
        "green": 13,
        "blue":  14,
    }
    var sum int

    file, err := os.Open("aoc.in")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        game := strings.Split(line, ": ")
        gameId, _ := strconv.Atoi(strings.Split(game[0], " ")[1])
        takes := strings.Split(game[1], "; ")
        isPossible := true
        for _, take := range takes {
            cubes := strings.Split(take, ", ")
            for _, cube := range cubes {
                number, color := parseCube(cube)
                if number > maxCubes[color] {
                    isPossible = false
                }
            }
        }
        if isPossible {
            sum = sum + gameId
        }
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    fmt.Println(sum)
}

