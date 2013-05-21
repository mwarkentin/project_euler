package main

import "fmt"

func prime_factors(n int) {
    d := 2
    for n > 1 {
        for n % d == 0 {
            fmt.Println("Prime factor:", d)
            n /= d
        }
        d++
    }
}

func main() {
    prime_factors(600851475143)
}
