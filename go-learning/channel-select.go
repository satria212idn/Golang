package main

import "fmt"
import "runtime"

func main() {
  runtime.GOMAXPROCS(2)

  var number = [] int{1,2,3,4,5,6,7,8,9,10}
  fmt.Println("Number :",number)

  var channel1 = make(chan float64)
  go get_avg(number, channel1)

  var channel2 = make(chan int)
  go value_max(number, channel2)

  for i:=0; i<2; i++ {
    select {
    case avg := <- channel1:
      fmt.Printf("Average \t : %.2f \n", avg)
    case max := <-channel2:
      fmt.Printf("Maximal \t : %d \n", max)
    }
  }
}

func get_avg(number [] int, ch chan float64)  {
  var sum = 0
  for _, e := range number {
    sum += e
  }
  ch <- float64(sum) / float64(len(number))
}

func value_max(number [] int, ch chan int) {
  var max = number[0]
  for _, e := range number {
    if max < e {
      max = e
    }
  }
    ch <- max
}
