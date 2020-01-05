package main

import "fmt"

func main() {
  var x student
  var x1 = student{"Mas Fian 77",19}
  x.name = "Alfian Satria Vivaldi Payu"
  x.age = 19
  x.myname()
  x1.myname()
}

type student struct {
  name string
  age int
}

func (y student) myname() {
  fmt.Println("My name is",y.name)
  fmt.Println("Age",y.age)
}
