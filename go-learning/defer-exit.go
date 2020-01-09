package main

import "fmt"
import "os"

func main() {
  defer fmt.Println("Hello")
  fmt.Println("Welcome")
  fmt.Println("To My Home")
  show()
  os.Exit(1)
}

func show() {
  fmt.Println("Show Me")
}
