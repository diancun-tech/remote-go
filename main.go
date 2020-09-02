package main

import "fmt"

func sum(a, b float64) float64 {
    return a + b
}

func avg(a, b float64) float64 {
    return sum(a, b) / 2
}

func main() {
    a := 3.1415926
    b := 1.5
    c := avg(a, b)
    fmt.Printf("Hello World, %f", c)
 }

