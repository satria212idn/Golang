package main

import "fmt"
import "reflect"

func main() {
  var number = 27
  var reflectnumber = reflect.ValueOf(number)

  fmt.Println("Type Variable :",reflectnumber.Type())

  if reflectnumber.Kind() == reflect.Int {
    fmt.Println("Value Variable :",reflectnumber.Int())
  }
}
