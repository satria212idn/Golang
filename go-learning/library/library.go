package library

import "fmt"

func isprivate() {
  fmt.Println("This is private")
}

func Ispublic() {
  fmt.Println("This is public")
  isprivate()
}
