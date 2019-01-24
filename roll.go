package main

import (
    "fmt"
    "math/rand"
    "time"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func main() {
  fmt.Println(`=====================================
Welcome to dice roller app!
Are you ready to roll the dice? (Y/N)
=====================================`)

  reader := bufio.NewReader(os.Stdin)
  resp, err := reader.ReadString('\n')

  if err != nil {
    fmt.Println(err)
    return
  }

  resp = strings.Replace(resp, "\n", "", -1)

  switch resp {
  case "Y":
    roll()
    fmt.Println(`Thanks, see you later!
=====================================`)
    break
  case "N":
    fmt.Println(`=====================================
Okay thanks, we're ready when you are!
=====================================`)
    break
  default:
    return
  }
}

func roll()  {
  fmt.Println(`=====================================
Okay! rolling...`)
  rand.Seed(time.Now().UTC().UnixNano())
  result := strconv.Itoa(randomizer(10))
  fmt.Println("You got " + result)
}

func randomizer(int) int{
    result := randInt(1, 6)
    return int(result)
}

func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}
