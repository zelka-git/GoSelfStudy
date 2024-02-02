package main

import "fmt"

func main(){
    var sum int = 0

label:
    for i:=0; i<10; i++ {
    for j:= 0; j < 11 ; j++ {
        break label;
    }
        if i%2 == 1 {
            fmt.Println(i, "is odd")
        } else {
        fmt.Println(i, "is even")
        }

        sum +=i
    }
    fmt.Println("sum is", sum)
}