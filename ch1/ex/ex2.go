package main // Package main is special. It defines a standalone 
             // executable program, not a library. 

import (
    "fmt"
    "os"
)


/**
  Modify the `echo` program to print the index and value
  of each of its arguments, one per line.
*/
func main() {
    for idx, arg := range os.Args[1:] {
        fmt.Println(idx, arg)
    }
}


func _main() {
    var s, sep string
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}


