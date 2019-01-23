package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    fmt.Println(randomString(10))
}

func randomString(int) int{
    result := randInt(1, 6)
    //fmt.Println(result)
    return int(result)
}

func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}
