package main

import "fmt"

func main() {
  var fruit = [] string{"Apple","Mango","Orange","Grape"}
  fruit = append(fruit,"Durian","Avocado") // this is slice
  fmt.Println(fruit)
  fmt.Println("Total Element :",len(fruit))
}
