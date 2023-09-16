package main

import "fmt"

func plus() {
    a := 32
    b := 65
    fmt.Println(a+b)
}

func whoami() {
    name := "김성연"
    fmt.Printf("나는 %s 입니다.\n", name)
}

func main() {
    plus()
    whoami()
}