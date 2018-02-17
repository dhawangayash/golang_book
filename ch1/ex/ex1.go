package main // Package main is special. It defines a standalone 
             // executable program, not a library. 


import (
    "fmt"
    "os"
)

func main() {
    var s, sep string
    for i := 0; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }
    fmt.Println(s)
}


