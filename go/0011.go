package main

import "fmt"
import "io/ioutil"
import "strconv"
import "strings"

func main() {
    var grid = [20][20]int {}
    greatest := 0

    content, err := ioutil.ReadFile("0011.txt")
    if err != nil {
        fmt.Println(content)
    }
    lines := strings.Split(string(content), "\n")
    for i := 0; i < 20; i++ {
        nums := strings.Split(lines[i], " ")
        for j := 0; j < 20; j++ {
            grid[i][j], err = strconv.Atoi(nums[j])
        }
    }

    // Get horizontal products
    for i := 0; i < 20; i++ {
        for j := 0; j < 16; j++ {
            product := grid[i][j]*grid[i][j+1]*grid[i][j+2]*grid[i][j+3]
            if product > greatest {
                greatest = product
            }
        }
    }

    // Get vertical product
    for i := 0; i < 16; i++ {
        for j := 0; j < 20; j++ {
            product := grid[i][j]*grid[i+1][j]*grid[i+2][j]*grid[i+3][j]
            if product > greatest {
                greatest = product
            }
        }
    }

    // Down + right
    for i := 0; i < 16; i++ {
        for j := 0; j < 16; j++ {
            product := grid[i][j]*grid[i+1][j+1]*grid[i+2][j+2]*grid[i+3][j+3]
            if product > greatest {
                greatest = product
            }
        }
    }

    // Up + right
    for i := 3; i < 20; i++ {
        for j := 0; j < 16; j++ {
            product := grid[i][j]*grid[i-1][j+1]*grid[i-2][j+2]*grid[i-3][j+3]
            if product > greatest {
                greatest = product
            }
        }
    }

    fmt.Println(greatest)
}
