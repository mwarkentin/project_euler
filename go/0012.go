package main

import "fmt"
import "os"

func main() {
    sum := 0
    for i := 1; ; i++ {
        numFactors := 0
        sum += i

        for j := 1; j <= sum; j++ {
            if sum % j == 0 {
                numFactors++
            }
        }

        fmt.Println(i, sum, numFactors)

        if numFactors > 500 {
            os.Exit(0)
        }
    }
}
