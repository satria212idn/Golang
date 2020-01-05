package main

import "fmt"

func main() {
  var number int = 27
  var address_number *int = &number // this is pointer

  fmt.Println("Value of number :",number)
  fmt.Println("Addres of number :",address_number)
}
