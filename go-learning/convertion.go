package main

import "fmt"
import "strconv"

func main() {
  var num = 27.7
  var str = strconv.FormatFloat(num,'f',6,64)

  fmt.Println(str)
  fmt.Println(num + 10)
}
