package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    for1();
    for2();
    for3();
    fmt.Println(strings.Join(os.Args[1:], " "))
    fmt.Println(os.Args[1:])
}

func for1(){
    var s, step string
    for i :=1; i < len(os.Args); i++ {
        s += step + os.Args[i]
        step = " "
    }
    fmt.Println(s)
}

func for2() {
    var s, step string
    i:=1
    for i < len(os.Args){
        s += step + os.Args[i]
        step = " "
        i++
    }
     fmt.Println(s)
}

func for3() {
    var s, step string
    for _, arg := range os.Args[1:] {
        s += step + arg
        step = " "
    }
    fmt.Println(s)
}