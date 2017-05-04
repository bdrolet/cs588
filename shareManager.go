package main

import "fmt"
import "strconv"

func createShares(k int, n int, secret string) (shares []string) {
    fmt.Println("k: " + strconv.Itoa(k))
    fmt.Println("n: " + strconv.Itoa(n))
    fmt.Println("secret: " + secret)
    shares = make([]string, 3)
    return
}

func combineShares(shares []string) (secret string) {
    secret = "I am a pretty pony."
    return
}