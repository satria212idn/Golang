package main

import "fmt"
import "encoding/base64"

func main() {
  var my_name = "MAS FIAN 77"

  var encodeString = base64.StdEncoding.EncodeToString([]byte(my_name))
  fmt.Println("Encoding :", encodeString)

  var decodedByte,_ = base64.StdEncoding.DecodeString("TUFTIEZJQU4gNzc=")
  var decodedString = string(decodedByte)
  fmt.Println("Decoding :", decodedString)
}
