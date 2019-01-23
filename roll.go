package main

import (
    "fmt"
    "math/rand"
    "time"
    "bufio"
    "os"
    "strings"
)

func main() {
  fmt.Print(
    `=====================================
    Welcome to dice roller app!
    Are you ready to roll the dice? (Y/N)
    =====================================`
  )

  reader := bufio.NewReader(os.Stdin)
  resp, _, err := reader.ReadRune()

  if err != nil {
    fmt.Println(err)
    return
  }

  switch resp {
  case "Y":
    // call roll()
    break
  case "N"
    // call okay we are ready when you are
    break
  }
}

func roll()  {
  rand.Seed(time.Now().UTC().UnixNano())
  fmt.Println(randomString(10))
}

func randomString(int) int{
    result := randInt(1, 6)
    return int(result)
}

func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}
