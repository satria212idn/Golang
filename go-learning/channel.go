package main

import "fmt"
import "runtime"

var msg = make(chan string)
func main() {
  runtime.GOMAXPROCS(2)

  go hello("P")
  go hello("I")
  go hello("V")

  var message1 = <-msg
  fmt.Println(message1)

  var message2 = <-msg
  fmt.Println(message2)

  var message3 = <-msg
  fmt.Println(message3)

}

func hello(name string) {
  var data = fmt.Sprintf("Hello %s", name)
  msg <- data
}
