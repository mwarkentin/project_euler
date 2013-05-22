package main

import "fmt"
import "strconv"

func isPalindromic(product int) bool {
    s := strconv.Itoa(product)
    n := len(s)
    runes := make([]rune, n)

    for i, rune := range s {
        runes[i] = rune
    }

    s1 := runes[:n/2]
    s2 := runes[n/2:]
    for i, _ := range s1 {
        if s1[i] != s2[len(s2) - 1 - i] {
            return false
        }
    }
    return true
}

func main() {
    n1 := 999
    n2 := 999
    largestPalindrome := 0

    for n1 >= 100 {
        for n2 >= 100 {
            product := n1 * n2
            if isPalindromic(product) {
                if product > largestPalindrome {
                    largestPalindrome = product
                }
            }
            n2--
        }
        n1--
        n2 = 999
    }
    fmt.Println(largestPalindrome)
}
