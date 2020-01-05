package main

import "fmt"

func main() {
  var x student
  x.name = "MasFian"
  x.class = 77
  x.age = 19

  fmt.Println("Name :",x.name)
  fmt.Println("Class :",x.class)
  fmt.Println("Age :",x.age)
}

type student struct {
  name string
  class int
  age int
}
