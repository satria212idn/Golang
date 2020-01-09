package main

import "fmt"
import "time"

func main() {
  fmt.Println("Start")
  time.Sleep(time.Second * 3)
  fmt.Println("Done after 3 second")
}
