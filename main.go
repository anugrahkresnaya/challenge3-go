package main

import "fmt"

func main()  {
  var greeting string = "selamat malam"
  for i := 0; i < len(greeting); i++  {
    fmt.Printf("%c\n", greeting[i])
  }
  mapGreeting := map[string]int{
    "s": 1,
    "e": 1,
    "l": 2,
    "a": 4,
    "m": 3,
    "t": 1,
    " ": 1,
  }
  fmt.Println(mapGreeting)
}
