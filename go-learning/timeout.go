package main

import "fmt"
import "runtime"
import "math/rand"
import "time"

func main() {
  rand.Seed(time.Now().Unix())
  runtime.GOMAXPROCS(2)

  var message = make(chan int)

  go send_data(message)
  receive_data(message)
}

func send_data(ch chan <- int) {
  for i:= 0; true; i++{
    ch <- i
    time.Sleep(time.Duration(rand.Int() %10+1) * time.Second)
  }
}

func receive_data(ch <- chan int)  {
  loop:
  for{
    select{
    case data := <- ch:
      fmt.Print("Receive Data ",data,"\n")
    case <- time.After(time.Second * 5):
      fmt.Println("Request Timeout")
      break loop
    }
  }
}
