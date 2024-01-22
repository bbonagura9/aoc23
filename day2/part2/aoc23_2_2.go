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
    var sum int

    file, err := os.Open("aoc.in")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        var minCubes = map[string]int {
            "red":   0,
            "green": 0,
            "blue":  0,
        }

        line := scanner.Text()
        game := strings.Split(line, ": ")
        takes := strings.Split(game[1], "; ")
        for _, take := range takes {
            cubes := strings.Split(take, ", ")
            for _, cube := range cubes {
                number, color := parseCube(cube)
                if number > minCubes[color] {
                    minCubes[color] = number
                }
            }
        }
        gamePower := minCubes["red"] * minCubes["green"] * minCubes["blue"]
        sum = sum + gamePower
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    fmt.Println(sum)
}

