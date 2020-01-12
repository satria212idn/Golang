package main

import "fmt"
import "net/url"

func main() {
  var url_text = "http://google.com:8888/library/?find=MF77 Channel &author=MAS FIAN 77 &year=2020"
  var u, err = url.Parse(url_text)

  if err != nil {
    fmt.Println(err.Error())
    return
  }

  fmt.Println("URI :", url_text)
  fmt.Println("Host :", u.Host)
  fmt.Println("Path :", u.Path)

  var find = u.Query()["find"][0]
  var author = u.Query()["author"][0]
  var year = u.Query()["year"][0]

  fmt.Println("Find :", find)
  fmt.Println("Author :", author)
  fmt.Println("Year :", year)
}
