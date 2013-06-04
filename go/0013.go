package main

import "fmt"
import "io/ioutil"
import "math/big"
import "strings"

func main() {
    sum := new(big.Int)
    content, err := ioutil.ReadFile("0013.txt")
    if err != nil {
        fmt.Println(content)
    }
    lines := strings.Split(string(content), "\n")

    for i := 0; i < 100; i++ {
        num := new(big.Int)
        num.SetString(lines[i], 10)
        sum.Add(sum, num)
    }
    fmt.Println(sum.String()[:10])
}
