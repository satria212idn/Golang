package main

import "fmt"

func main() {
  fmt.Println(show_message(10,5))
}

func show_message(x int, y int)int {
  var z = x * y
  return z
}
