package main

import "fmt"
import "runtime"

func main() {
  runtime.GOMAXPROCS(2)

  msg := make(chan int, 2)

  go func() {
    for {
      i := <- msg
      fmt.Println("Receive data",i)
    }
  }()

  for i:=0; i<5; i++ {
    fmt.Println("Send data",i)
    msg <- i
  }
}
