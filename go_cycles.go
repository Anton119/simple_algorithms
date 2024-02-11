package main 

import "fmt"

func main () {
    for i:=0; i<=4; i++ {
        for j:=0; j<=i; j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
fmt.Println()

for i:=0; i<=4; i++ {
    for j:=1; j<(5-i); j++ {
        fmt.Print(" ")
    }
    for k:=0; k<=i; k++ {
        fmt.Print("*")
    }
        
    fmt.Println()
}
fmt.Println()

for i:=0; i<=4; i++ {
    for j:=1; j<=(5-i); j++ {
        fmt.Print(" ")
    }
    for k:=1; k<=i; k++ {
        fmt.Print("* ")
    }
    fmt.Println()
}
}