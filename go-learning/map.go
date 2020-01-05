package main

import "fmt"

func main() {
  var price = map[string]int{"fried_chicken":10000, "fried_potato":8000}
  fmt.Println("Price of : \n ===================")
  fmt.Println("Fried Chicken :",price["fried_chicken"])
  fmt.Println("and")
  fmt.Println("Fried Potato :",price["fried_potato"])
}
