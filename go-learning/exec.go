package main

import "fmt"
import "os/exec"

func main() {
  var result,_ = exec.Command("ifconfig").Output()
  fmt.Println(string(result))
}
