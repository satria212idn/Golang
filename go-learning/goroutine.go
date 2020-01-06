package main

import "fmt"
import "runtime"

func main() {
  runtime.GOMAXPROCS(2)

  go show_msg(5, "Sending")
  show_msg(5, "Sent")

  var input string
  fmt.Scanln(&input)
}

func show_msg(x int, msg string) {
  for i := 0; i<x; i++ {
    fmt.Println((i +1), msg)
  }
}
