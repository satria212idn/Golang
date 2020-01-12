package main

import "fmt"
import "encoding/json"
import "net/http"
import "bytes"
import "net/url"

func main() {
  var menu, err = get_api("Nasi Goreng")
  if err != nil {
    fmt.Println("Sorry, menu not available", err.Error())
    return
  }
    fmt.Println("ID :", menu.ID, "Name :", menu.Name, "Price :", menu.Price)
}

var baseURL = "http://localhost:8080"

type food_menu struct {
  ID string
  Name string
  Price int
}

func get_api(menu string)(food_menu, error) {
  var err error
  var client = &http.Client{}
  var data food_menu

  var param = url.Values{}
  param.Set("Name", menu)

  var payload = bytes.NewBufferString(param.Encode())

  request, err := http.NewRequest("POST", baseURL + "/find_menu", payload)

  if err != nil {
    return data, err
  }

  request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  response, err := client.Do(request)

  if err != nil {
    return data, err
  }
  defer response.Body.Close()

  err = json.NewDecoder(response.Body).Decode(&data)
  if err != nil {
    return data, err
  }
  return data, nil
}
