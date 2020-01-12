package main

import "fmt"
import "net/http"

func index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Welcome, MAS FIAN 77!")
}

func main() {
  http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w, "Hello World!")
  })
  http.HandleFunc("/index", index)
  fmt.Println("Starting web server on localhost:8080")
  http.ListenAndServe(":8080",nil)
}
