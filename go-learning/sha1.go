package main

import "crypto/sha1"
import "fmt"

func main() {
  var text = "This is Secret Zone"
  var sha = sha1.New()
  sha.Write([]byte(text))
  var encrypt = sha.Sum(nil)
  var stringencrypt = fmt.Sprintf("%x",encrypt)

  fmt.Println(stringencrypt)

}
