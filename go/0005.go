package main

import "fmt"
import "os"

func main() {
    for i := 2520; ; i++ {
        for j := 1; j <= 20; j++ {
            if i % j != 0 {
                break
            }
            if j == 20 {
                fmt.Println(i)
                os.Exit(0)
            }
        }
    }
}
