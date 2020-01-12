package main

import "fmt"
import "database/sql"
import _ "mysql-master"

type student struct {
  id int
  name string
  age int
}

func connection() (*sql.DB, error) {
  db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_golang")
  if err != nil {
    return nil, err
  }
  return db, nil
}

// CREATE

func sql_create()  {
  db, err := connection()
  if err != nil {
    fmt.Println(err.Error())
    return
  }
  defer db.Close()

  _, err = db.Exec("INSERT INTO student VALUES (?,?,?)",nil,"GO Language",12)

  if err != nil {
    fmt.Println(err.Error())
  }
  fmt.Println("Success.")
}

// READ

func sql_read() {
  db, err := connection()
  if err != nil {
    fmt.Println(err.Error())
    return
  }
  defer db.Close()

  rows , err := db.Query("SELECT * FROM student")

  if err != nil {
    fmt.Println(err.Error())
    return
    }
    defer rows.Close()

    var result []student

    for rows.Next() {
      var each = student{}

      var err = rows.Scan(&each.id, &each.name, &each.age)

      if err != nil {
        fmt.Println(err.Error())
        return
      }
      result = append(result, each)
    }
    if err = rows.Err(); err != nil {
      fmt.Println(err.Error())
      return
    }
    for _, each := range result {
      fmt.Println(each.id, each.name, each.age)
      }
}

// UPDATE

func sql_update() {
  db, err := connection()
  if err != nil {
    fmt.Println(err.Error())
    return
  }
  defer db.Close()

  _, err = db.Exec("UPDATE student SET name = ? WHERE age = ?", "MAS FIAN", 19)
  if err != nil {
    fmt.Println(err.Error())
    return
  }
  fmt.Println("Update success.")
}

// DELETE

func sql_delete()  {
  db, err := connection()
    if err != nil {
      fmt.Println(err.Error())
      return
    }
    defer db.Close()

    _, err = db.Exec("DELETE FROM student WHERE name = ?", "MAS FIAN")
    if err != nil {
      fmt.Println(err.Error())
      return
    }
    fmt.Println("Delete success.")
  }


func main() {
  //sql_create()
  //sql_update();
  //sql_delete()
  sql_read() // always last !!!
}
