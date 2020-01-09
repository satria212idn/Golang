package main

import "fmt"
import "regexp"

func main() {
  var text = "Orange BANAN4 GRAPE"
  var regex, err = regexp.Compile(`[A-Z]+\d`)

  if err != nil {
    fmt.Println(err.Error())
  }

  var result = regex.FindAllString(text, -1)
  fmt.Println(result)

  var match = regex.MatchString(text)
  fmt.Println(match)

  var index1 = regex.FindStringIndex(text)
  fmt.Println(index1)

  var new_string = regex.ReplaceAllString(text, "DURIAN")
  fmt.Println(new_string)
}
