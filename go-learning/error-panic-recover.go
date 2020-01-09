package main

import "fmt"
import "strconv"

func main() {
  defer recover_me()
  var input string
  fmt.Print("Input a number : ")
  fmt.Scanln(&input)

  var number int
  var err error
  number, err = strconv.Atoi(input)

  if err == nil {
    fmt.Println(number, "This is a number")
  } else {
    fmt.Println(input, "This not a number")
    panic(err.Error())
    fmt.Println("Show Me")
  }
}

func recover_me() {
  if r:=recover(); r!=nil {
    fmt.Println("Finally showed")
  } else {
    fmt.Println("It's OK")
  }
}
