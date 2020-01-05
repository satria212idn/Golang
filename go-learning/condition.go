package main

import "fmt"

func main() {
  var nilai = 10

  switch nilai {
  case 10:
    fmt.Println("Excellent")
  case 9:
    fmt.Println("Good")
  case 8:
    fmt.Println("Not Bad")
  default:
    fmt.Println("Try Again")
  }

  // if nilai == 10 {
  //   fmt.Println("Excellent")
  // } else if nilai >= 7 {
  //   fmt.Println("Good")
  // } else if nilai == 6 {
  //   fmt.Println("Not Bad")
  // } else {
  //   fmt.Println("Try Again")
  // }
}
