package main

import "fmt"

func main() {
    total := 0

    oldFirst := 0
    first := 0
    second := 1

    for second < 4000000 {
        oldFirst = first
        first = second
        second += oldFirst

        if second % 2 == 0 {
            total += second
        }
    }
    fmt.Println(total)
}
