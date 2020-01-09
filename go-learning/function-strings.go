package main

import "fmt"
import "strings"

func main() {
  var text = "Apple Fruit"
  var search = "Apple"
  var new_text = strings.Replace(text,search,"Orange",1)

  fmt.Println(text)
  fmt.Println(new_text)
}
