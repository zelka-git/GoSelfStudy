package main

import (
    "bufio"
    "fmt"
    "os"
)

func main(){
var input = readLine("Enter yes or no: ")

switch {
    case input == "yes" || input =="да":
        fmt.Println("agree")
        fallthrough //чтобы выполнелось следующее тоже
    case input == "no"|| input =="нет":
        fmt.Println("disagree")
    default:
        fmt.Println("Opps!")
    }

}

func readLine(line string) string {
    fmt.Println(line)
    reader := bufio.NewReader(os.Stdin)
    text, _,_ := reader.ReadLine()
    return string(text)
}