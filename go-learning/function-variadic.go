package main

import "fmt"

func main() {
  var average = count(1,2,3,4,5,6,7,8,9,10)
  var message = fmt.Sprintf("Rata-rata : %.2f", average)
  fmt.Println(message)
}

func count(number ...int)float64 {
  var total int = 0
  for _, number := range number {
    total += number
    fmt.Println(number)
  }
  var avg = float64(total) / float64(len(number))
  return avg
}
