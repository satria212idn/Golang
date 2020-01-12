package main

import "fmt"
import "encoding/json"
import "net/http"
import "database/sql"
import _ "mysql-master"

func main() {
  get_menu_db()
  http.HandleFunc("/menu",get_menu)
  http.HandleFunc("/find_menu", find_menu)
  fmt.Println("Starting web server on localhost:8080")
  http.ListenAndServe(":8080",nil)
}

type food_menu struct {
  ID string
  Name string
  Price int
}

func connection() (*sql.DB, error) {
  db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_menu")

  if err != nil {
    return nil, err
  }
  return db, nil
}
// POST data by POSTMAN without Database

// var data = []food_menu {
//   food_menu{"F01","Fried Rice",13000},
//   food_menu{"F02","Fried Chicken",12000},
//   food_menu{"F03","Fried Potato",11000},
//   food_menu{"F04","Fried Duck",10000},
//   food_menu{"D01","Avocado Juice",9000},
//   food_menu{"D02","Mango Juice",8000},
// }

func get_menu(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  if r.Method == "POST" {
    var result, err = json.Marshal(data)

    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    w.Write(result)
    return
  }
  http.Error(w, "", http.StatusBadRequest)
}

// GET data with Database

var data []food_menu

func get_menu_db()  {
  db, err := connection()

  if err != nil {
    fmt.Println(err.Error())
    return
  }
  defer db.Close()

  rows, err := db.Query("SELECT * FROM food_menu")
  if err != nil {
    fmt.Println(err.Error())
    return
  }
  defer rows.Close()

  for rows.Next() {
    var each = food_menu{}
    var err = rows.Scan(&each.ID, &each.Name, &each.Price)

    if err != nil {
      fmt.Println(err.Error())
      return
    }
    data = append(data, each)
  }
  if err = rows.Err(); err != nil {
    fmt.Println(err.Error())
    return
  }
}

func find_menu(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  if r.Method == "POST" {
    var name = r.FormValue("Name")
    var result []byte
    var err error

    for _, each := range data {
      if each.Name == name {
        result, err = json.Marshal(each)

        if err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
        }
        w.Write(result)
        return
      }
    }
    http.Error(w, "Sorry, menu not available", http.StatusBadRequest)
    return
  }
  http.Error(w, "", http.StatusBadRequest)
}
