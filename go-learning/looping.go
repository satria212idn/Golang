package main

import "fmt"

func main() {
  // //for
  // for i := 0; i<11; i++ {
  //   fmt.Println("Number",i)
  // }

  // while
  // var i=0
  // for i<11 {
  //   fmt.Println("Number",i)
  //   i++
  // }

  //for without argument
  var i=0
  for {
    fmt.Println("Number",i)
    i++ // to make infinite you can add comment if below
    if i==11 {
      break
    }
  }
}
