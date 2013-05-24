// Euclid's method: http://logicville.com/sel1.htm
// x = 10
// y = 5

// A = 100 - 25 = 75
// B = 2 * 10 * 5 = 100
// C = 100 + 25 = 125

// A*A + B*B = 15625
// C*C = 15625

package main

import "fmt"

func main() {
    for x := 5; x < 999; x++ {
        for y := 4; y < x; y++ {
            a := x*x - y*y
            b := 2*x*y
            c := x*x + y*y
            if a + b + c == 1000 {
                fmt.Println("a", a, "b", b, "c", c, "product:", a*b*c)
            }
        }
    }
}
