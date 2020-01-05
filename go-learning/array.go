package main

import "fmt"

func main() {
  var fruit=[3]string{"Apple","Orange","Grape"}
  fruit[1]="Mango"
  fmt.Println("Total Element :",len(fruit))
  fmt.Println("Content :",fruit)
}
