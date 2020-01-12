package main

import "fmt"
import "encoding/json"

type User struct {
  Name string `json:"Name"`
  Age int
}

 // decode json to object

// func main() {
//   var jsonString = `{"Name" : "MAS FIAN 77", "Age" : 19}`
//   var jsonData = []byte(jsonString)
//
//   var data User
//
//   var err = json.Unmarshal(jsonData, &data)
//   if err != nil {
//     fmt.Println(err.Error())
//     return
//   }
//
//   fmt.Println("Name :", data.Name)
//   fmt.Println("Age :", data.Age)
// }

// encode object to json

func main() {
  var object = []User{{"MAS FIAN 77", 19}, {"MF77", 19}}
  var jsonData, err = json.Marshal(object)

  if err != nil {
    fmt.Println(err.Error())
    return
  }

  var jsonString = string(jsonData)
  fmt.Println(jsonString)
}
